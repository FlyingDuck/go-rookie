package httpd

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Request struct {
	Method     string
	URL        *url.URL
	Proto      string
	Header     Header
	Body       io.Reader
	RemoteAddr string
	RequestURI string //字符串形式的url

	//私有字段
	conn           *conn
	cookies        map[string]string
	queryString    map[string]string
	contentType    string
	boundary       string
	postForm       map[string]string
	multipartForm  *MultipartForm
	haveParsedForm bool
	parseFormErr   error
}

func (r *Request) Query(key string) string {
	return r.queryString[key]
}

func (r *Request) Cookie(key string) string {
	if r.cookies == nil {
		r.parseCookie()
	}
	return r.cookies[key]
}

func (r *Request) FormFile(key string)(fh* FileHeader,err error){
	mf,err := r.MultipartForm()
	if err!=nil{
		return
	}
	fh,ok:=mf.File[key]
	if !ok{
		return nil,errors.New("http: missing multipart file")
	}
	return
}

func (r *Request) parseCookie() {
	if r.cookies != nil {
		return
	}
	r.cookies = make(map[string]string)
	rawCookies, ok := r.Header["Cookie"]
	if !ok {
		return
	}
	for _, line := range rawCookies {
		kvs := strings.Split(strings.TrimSpace(line), ";")
		if len(kvs) == 1 && kvs[0] == "" {
			continue
		}
		for i := 0; i < len(kvs); i++ {
			index := strings.IndexByte(kvs[i], '=')
			if index == -1 {
				continue
			}
			key := strings.TrimSpace(kvs[i][:index])
			value := strings.TrimSpace(kvs[i][index+1:])
			r.cookies[key] = value
		}
	}
}

func (r *Request) setupBody() {
	if r.Method != "POST" {
		r.Body = new(eofReader)
		return
	}
	if r.chunked() {
		r.Body = newChunkReader(r.conn.bufr)
		r.fixExpectContinueReader()
		return
	}

	contentLen, err := strconv.ParseInt(r.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		r.Body = new(eofReader)
		return
	}
	r.Body = io.LimitReader(r.conn.bufr, contentLen)
	r.fixExpectContinueReader()
}

func (r *Request) chunked() bool {
	return r.Header.Get("Transfer-Encoding") == "chunked"
}

// 些客户端在发送完 http 首部之后，发送body数据前，会先通过发送Expect: 100-continue查询服务端是否希望接受body数据，
// 服务端只有回复了HTTP/1.1 100 Continue客户端才会再次发送body。
func (r *Request) fixExpectContinueReader() {
	if r.Header.Get("Expect") != "100-continue" {
		return
	}
	r.Body = &expectContinueReader{
		r: r.Body,
		w: r.conn.bufw,
	}
}

// boundary是存取在Content-Type字段中
func (r *Request) parseContentType() {
	ct := r.Header.Get("Content-Type")
	//Content-Type: multipart/form-data; boundary=------974767299852498929531610575
	//Content-Type: multipart/form-data; boundary=""------974767299852498929531610575"
	//Content-Type: application/x-www-form-urlencoded
	index := strings.IndexByte(ct, ';')
	if index == -1 {
		r.contentType = ct
		return
	}
	if index == len(ct)-1 {
		return
	}
	ss := strings.Split(ct[index+1:], "=")
	if len(ss) < 2 || strings.TrimSpace(ss[0]) != "boundary" {
		return
	}
	// 将解析到的CT和boundary保存在Request中
	r.contentType, r.boundary = ct[:index], strings.Trim(ss[1], `"`)
	return
}

// 得到一个MultipartReader
func (r *Request) MultipartReader() (*MultipartReader, error) {
	if r.boundary == "" {
		return nil, errors.New("no boundary detected")
	}
	return NewMultipartReader(r.Body, r.boundary), nil
}

// 如果用户在Handler的回调函数中没有去读取Body的数据，就意味着处理同一个 socket 连接上的下一个http报文时，
// Body未消费的数据会干扰下一个http报文的解析。所以我们的框架还需要在 Handler 结束后，将当前http请求的数据给消费掉。
func (r *Request) finish() error {
	// 将缓存中的剩余的数据发送到 rwc 中
	if err := r.conn.bufw.Flush(); err != nil {
		return err
	}
	// 消费掉剩余的数据
	_, err := io.Copy(io.Discard, r.Body)
	return err
}

func (r *Request) PostForm(name string) string {
	if !r.haveParsedForm {
		r.parseFormErr = r.parseForm()
	}
	if r.parseFormErr != nil || r.postForm == nil {
		return ""
	}
	return r.postForm[name]
}

func (r *Request) MultipartForm() (*MultipartForm, error) {
	if !r.haveParsedForm {
		if err := r.parseForm(); err != nil {
			r.parseFormErr = err
			return nil, err
		}
	}
	return r.multipartForm, r.parseFormErr
}

func (r *Request) parseForm() error {
	if r.Method != "POST" && r.Method != "PUT" {
		return errors.New("missing form body")
	}
	r.haveParsedForm = true
	switch r.contentType {
	case "application/x-www-form-urlencoded":
		return r.parsePostForm()
	case "multipart/form-data":
		return r.parseMultipartForm()
	default:
		return errors.New("unsupported form type")
	}
}

func (r *Request) parsePostForm() error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	r.postForm = parseQuery(string(data))
	return nil
}

func (r *Request) parseMultipartForm() (err error) {
	mr,err := r.MultipartReader()
	if err!=nil{
		return
	}
	r.multipartForm,err = mr.ReadForm()
	//让PostForm方法也可以访问multipart表单的文本数据
	r.postForm = r.multipartForm.Value
	return
}

func (r *Request) finishRequest(resp *Response) (err error) {
	if r.multipartForm != nil {
		r.multipartForm.RemoveAll()
	}
	//告诉chunkWriter handler已经结束
	resp.handlerDone = true
	//触发chunkWriter的Write方法，Write方法通过handlerDone来决定是用chunk还是Content-Length
	if err = resp.bufw.Flush(); err != nil {
		return
	}
	//如果是使用chunk编码，还需要将结束标识符传输
	if resp.chunking{
		_,err = resp.c.bufw.WriteString("0\r\n\r\n")
		if err!=nil{
			return
		}
	}

	//如果用户的handler中未Write任何数据，我们手动触发(*chunkWriter).writeHeader
	if !resp.cw.wrote {
		resp.header.Set("Content-Length", "0")
		if err = resp.cw.writeHeader(); err != nil {
			return
		}
	}

	if err = r.conn.bufw.Flush(); err != nil {
		return
	}
	_, err = io.Copy(io.Discard, r.Body)
	return err
}

// readRequest 解析 HTTP 报文，如下：
// POST /index?name=gu HTTP/1.1\r\n			#请求行
// Content-Type: text/plain\r\n				#此处至报文主体为首部字段
// User-Agent: PostmanRuntime/7.28.0\r\n
// Host: 127.0.0.1:8080\r\n
// Accept-Encoding: gzip, deflate, br\r\n
// Connection: keep-alive\r\n
// Cookie: uuid=12314753; tid=1BDB9E9; HOME=1\r\n
// Content-Length: 18\r\n
// \r\n
// hello,I am client!							#报文主体
func readRequest(c *conn) (*Request, error) {
	r := new(Request)
	r.conn = c
	r.RemoteAddr = c.rwc.RemoteAddr().String()

	// 读取第一行
	line, err := readLine(c.bufr) // POST /index?name=gu HTTP/1.1
	if err != nil {
		return r, err
	}
	_, err = fmt.Sscanf(string(line), "%s%s%s", &r.Method, &r.RequestURI, &r.Proto)
	if err != nil {
		return nil, err
	}

	r.URL, err = url.ParseRequestURI(r.RequestURI)
	if err != nil {
		return nil, err
	}

	r.queryString = parseQuery(r.URL.RawQuery)
	r.Header, err = parseHeader(c.bufr)
	if err != nil {
		return nil, err
	}

	const noLimited = (1 << 63) - 1
	r.conn.lr.N = noLimited
	r.setupBody()

	r.parseContentType()
	return r, nil
}

func readLine(bufr *bufio.Reader) ([]byte, error) {
	line, isPrefix, err := bufr.ReadLine()
	if err != nil {
		return nil, err
	}

	var lineBuf bytes.Buffer
	lineBuf.Write(line)
	for isPrefix {
		line, isPrefix, err = bufr.ReadLine()
		if err != nil {
			return nil, err
		}
		lineBuf.Write(line)
	}
	return lineBuf.Bytes(), nil
}

func parseQuery(rawQuery string) (map[string]string) {
	parts := strings.Split(rawQuery, "&")
	queries := make(map[string]string, len(parts))

	for _, part := range parts {
		index := strings.IndexByte(part, '=')
		if index < 0 || index == len(part)-1 {
			continue
		}
		key := strings.TrimSpace(part[:index])
		value := strings.TrimSpace(part[index+1:])
		queries[key] = value
	}
	return queries
}

func parseHeader(bufr *bufio.Reader) (Header, error) {
	header := make(Header)
	for {
		line, err := readLine(bufr)
		if err != nil {
			return nil, err
		}
		//如果读到/r/n/r/n，代表报文首部的结束
		if len(line) == 0 {
			break
		}

		index := bytes.IndexByte(line, ':')
		if index == -1 {
			return nil, errors.New("unsupported protocol")
		}
		if index == len(line)-1 {
			continue
		}
		key := strings.TrimSpace(string(line)[:index])
		value := strings.TrimSpace(string(line)[index+1:])
		header[key] = append(header[key], value)
	}

	return header, nil
}

type eofReader struct {
}

func (r eofReader) Read([]byte) (int, error) {
	return 0, io.EOF
}

type expectContinueReader struct {
	wroteContinue bool // 是否已经发送过100 continue
	r             io.Reader
	w             *bufio.Writer
}

func (r *expectContinueReader) Read(p []byte) (n int, err error) {
	if !r.wroteContinue { // 第一次读取前发送100 continue
		r.w.WriteString("HTTP/1.1 100 Continue\r\n\r\n")
		r.w.Flush()
		r.wroteContinue = true
	}
	return r.r.Read(p)
}


type MultipartForm struct {
	Value map[string]string
	File  map[string]*FileHeader
}

type FileHeader struct {
	Filename string
	Header   Header
	Size     int
	content  []byte
	tmpFile  string
}

func (fh *FileHeader) Open() (io.ReadCloser, error) {
	if fh.inDisk() {
		return os.Open(fh.tmpFile)
	}
	b := bytes.NewReader(fh.content)
	return io.NopCloser(b), nil
}

func (fh *FileHeader) inDisk() bool {
	return fh.tmpFile != ""
}

func (fh *FileHeader) Save(dest string)(err error){
	rc,err:=fh.Open()
	if err!=nil{
		return
	}
	defer rc.Close()
	file,err:=os.Create(dest)
	if err!=nil{
		return
	}
	defer file.Close()
	_,err = io.Copy(file,rc)
	if err!=nil{
		os.Remove(dest)
	}
	return
}