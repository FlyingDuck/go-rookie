package httpd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const bufSize = 4096	// 滑动窗口大小

// MultipartReader
// POST /login HTTP/1.1\r\n
// [[ Less interesting headers ... ]]
// Content-Type: multipart/form-data; boundary=---------------------------735323031399963166993862150\r\n
// Content-Length: 414\r\n
// \r\n
// -----------------------------735323031399963166993862150\r\n		#--boundary，注意比上面的boundary多了两个-
// Content-Disposition: form-data; name="username"\r\n					#第一部分，username
// \r\n
// gu\r\n
// -----------------------------735323031399963166993862150\r\n		#--boundary
// Content-Disposition: form-data; name="password"\r\n					#第二部分，password
// \r\n
// 123\r\n
// -----------------------------735323031399963166993862150\r\n		#--boundary
// Content-Disposition: form-data; name="file1"; filename="1.txt"\r\n	#第三部分，文件1
// Content-Type: text/plain\r\n
// \r\n
// Content of 1.txt.\r\n
// -----------------------------735323031399963166993862150\r\n		#--boundary
// Content-Disposition: form-data; name="file2"; filename="2.html"\r\n	#第四部分，文件2
// Content-Type: text/html\r\n
// \r\n
// <!DOCTYPE html><title>Content of 2.html.</title>\r\n
// -----------------------------735323031399963166993862150--\r\n		#--bounadry--标识表单结束
type MultipartReader struct {
	// bufr 是对Body的封装，方便我们预查看Body上的数据，从而确定part之间边界
	// 每个part共享这个 bufr，但只有Body的读取指针指向哪个part的报文，
	// 哪个 part 才能在 bufr 上读取数据，此时其他 part 是无效的
	bufr                 *bufio.Reader
	// 记录bufr的读取过程中是否出现io.EOF错误，如果发生了这个错误，
	// 说明Body数据消费完毕，表单报文也消费完，不需要再产生下一个part
	occurEofErr          bool
	crlfDashBoundaryDash []byte				//\r\n--boundary--
	crlfDashBoundary     []byte				//\r\n--boundary，分隔符
	dashBoundary         []byte				//--boundary
	dashBoundaryDash     []byte				//--boundary--
	curPart              *Part				//当前解析到了哪个part
	crlf                 [2]byte			//用于消费掉\r\n
}

//传入的r将是Request的Body，boundary会在http首部解析时就得到
func NewMultipartReader(r io.Reader, boundary string) *MultipartReader {
	b := []byte("\r\n--" + boundary + "--")
	return &MultipartReader{
		bufr:                 bufio.NewReaderSize(r, bufSize),	//将io.Reader封装成bufio.Reader
		crlfDashBoundaryDash: b,
		crlfDashBoundary:     b[:len(b)-2],
		dashBoundary:         b[2 : len(b)-2],
		dashBoundaryDash:     b[2:],
	}
}

func (mr *MultipartReader) NextPart() (*Part, error) {
	if mr.curPart != nil {
		// 将当前的Part关闭掉，即消费掉当前part数据，好让body的读取指针指向下一个part
		if err := mr.curPart.Close(); err != nil {
			return nil, err
		}
		if err := mr.discardCRLF(); err != nil {
			return nil, err
		}
	}
	// 下一行就应该是boundary分割
	line, err := mr.readLine()
	if err != nil {
		return nil, err
	}
	// 到 multipart 报文的结尾了，直接返回
	if bytes.Equal(line, mr.dashBoundaryDash) {
		return nil, io.EOF
	}
	if !bytes.Equal(line, mr.dashBoundary) {
		err = fmt.Errorf("want delimiter %s, but got %s", mr.dashBoundary, line)
		return nil, err
	}
	// 这时Body已经指向了下一个part的报文
	p := new(Part)
	p.mr = mr
	// 前文讲到要将part的首部信息预解析，好让part指向消息主体，具体实现见后文
	if err = p.readHeader(); err != nil {
		return nil, err
	}
	mr.curPart = p
	return p, nil
}

// 消费掉\r\n
func (mr *MultipartReader) discardCRLF() (err error) {
	if _, err = io.ReadFull(mr.bufr, mr.crlf[:]); err == nil {
		if mr.crlf[0] != '\r' && mr.crlf[1] != '\n' {
			err = fmt.Errorf("expect crlf, but got %s", mr.crlf)
		}
	}
	return
}

// 读一行
func (mr *MultipartReader) readLine() ([]byte, error) {
	return readLine(mr.bufr)
}

// 直接利用了解析http报文首部的函数readHeader，很简单
func (p *Part) readHeader() (err error) {
	p.Header, err = parseHeader(p.mr.bufr)
	return err
}

// 将当前part剩余的数据消费掉，防止其报文残存在Reader上影响下一个part
func (p *Part) Close() error {
	if p.closed {
		return nil
	}
	_, err := io.Copy(io.Discard, p)
	p.closed = true	//标记状态为关闭
	return err
}


type Part struct {
	Header           Header			// 存取当前part的首部
	mr               *MultipartReader

	formName         string
	fileName         string			// 当该part传输文件时，fileName不为空
	closed           bool			// part是否关闭
	substituteReader io.Reader		// 替补Reader
	parsed           bool			// 是否已经解析过formName以及fileName
}

func (p *Part) Read(buf []byte) (n int, err error) {
	// part已经关闭后，直接返回io.EOF错误
	if p.closed {
		return 0, io.EOF
	}
	// 不为nil时，优先让substituteReader读取
	if p.substituteReader != nil {
		return p.substituteReader.Read(buf)
	}
	bufr := p.mr.bufr
	var peek []byte
	//如果已经出现EOF错误，说明Body没数据了，这时只需要关心bufr还剩余已缓存的数据
	if p.mr.occurEofErr {
		peek, _ = bufr.Peek(bufr.Buffered())	// 将最后缓存数据取出
	} else {
		//bufSize即bufr的缓存大小，强制触发Body的io，填满bufr缓存
		peek, err = bufr.Peek(bufSize)
		//出现EOF错误，代表Body数据读完了，我们利用递归跳转到另一个if分支
		if err == io.EOF {
			p.mr.occurEofErr = true
			return p.Read(buf)
		}
		if err != nil {
			return 0, err
		}
	}
	//在peek出的数据中找boundary
	index := bytes.Index(peek, p.mr.crlfDashBoundary)
	//两种情况：
	//1.即||前的条件，index!=-1 代表在peek出的数据中找到分隔符，也就代表顺利找到了该part的Read指针终点，
	//	给该part限制读取长度即可。
	//2.即||后的条件，在前文的multipart报文，是需要 boundary 来标识报文结尾，然后已经出现EOF错误,
	//  即在没有多余报文的情况下，还没有发现结尾标识，说明客户端没有将报文发送完整，就关闭了链接，
	//  这时让 substituteReader = io.LimitReader(-1)，逻辑上等价于 eofReader 即可
	if index != -1 || (index == -1 && p.mr.occurEofErr) {
		p.substituteReader = io.LimitReader(bufr, int64(index))
		return p.substituteReader.Read(buf)
	}

	//以下则是在peek出的数据中没有找到分隔符的情况，说明peek出的数据属于当前的part
	//见上文讲解，不能一次把所有的bufSize都当作消息主体读出，还需要减去分隔符的最长子串的长度。
	maxRead := bufSize - len(p.mr.crlfDashBoundary) + 1
	if maxRead > len(buf) {
		maxRead = len(buf)
	}
	return bufr.Read(buf[:maxRead])
}

// 获取FormName
func (p *Part) FormName() string {
	if !p.parsed {
		p.parseFormData()
	}
	return p.formName
}

// 获取FileName
func (p *Part) FileName() string {
	if !p.parsed {
		p.parseFormData()
	}
	return p.fileName
}

func (p *Part) parseFormData() {
	p.parsed = true
	cd := p.Header.Get("Content-Disposition")
	ss := strings.Split(cd, ";")
	if len(ss) == 1 || strings.ToLower(ss[0]) != "form-data" {
		return
	}
	for _, s := range ss {
		key, value := getKV(s)
		switch key {
		case "name":
			p.formName = value
		case "filename":
			p.fileName = value
		}
	}
}

func getKV(s string) (key string, value string) {
	ss := strings.Split(s, "=")
	if len(ss) != 2 {
		return
	}
	return strings.TrimSpace(ss[0]), strings.Trim(ss[1], `"`)
}