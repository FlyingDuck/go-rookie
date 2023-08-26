package httpd

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/url"
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
	r.Body, err = setupBody(c.bufr)
	if err != nil {
		return nil, err
	}
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

func setupBody(bufr *bufio.Reader) (io.Reader, error) {
	return new(eofReader), nil
}

type eofReader struct {
}

func (r eofReader) Read([]byte) (int, error) {
	return 0, io.EOF
}
