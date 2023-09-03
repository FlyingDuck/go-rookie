package httpd

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/url"
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
	conn        *conn
	cookies     map[string]string
	queryString map[string]string
	contentType string
	boundary    string
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
	r.contentType, r.boundary = ct[:index], strings.Trim(ss[1],`"`)
	return
}

// 得到一个MultipartReader
func (r *Request) MultipartReader()(*MultipartReader,error){
	if r.boundary==""{
		return nil,errors.New("no boundary detected")
	}
	return NewMultipartReader(r.Body,r.boundary),nil
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

	r.queryString, err = parseQuery(r.URL.RawQuery)
	if err != nil {
		return nil, err
	}
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

func parseQuery(rawQuery string) (map[string]string, error) {
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
	return queries, nil
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
