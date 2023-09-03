package httpd

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
)

func setupResponse(c *conn, req *Request) *Response {
	resp := &Response{
		c:          c,
		header:     make(Header),
		statusCode: 200,
		req:        req,
	}
	cw := &chunkWriter{resp: resp}
	resp.cw = cw
	resp.bufw = bufio.NewWriterSize(cw, 4096)
	var (
		protoMinor int
		protoMajor int
	)
	fmt.Sscanf(req.Proto, "HTTP/%d.%d", &protoMinor, &protoMajor)
	if protoMajor < 1 || protoMinor == 1 && protoMajor == 0 || req.Header.Get("Connection") == "close" {
		resp.closeAfterReply = true
	}
	return resp
}

type ResponseWriter interface {
	Write([]byte) (n int, err error)
	Header() Header
	WriteHeader(statusCode int)
}

type Response struct {
	//http连接
	c *conn

	//是否已经调用过WriteHeader，防止重复调用
	wroteHeader bool
	header      Header

	//WriteHeader传入的状态码，默认为200
	statusCode int

	//如果handler已经结束并且Write的长度未超过最大写缓存量，我们给头部自动设置Content-Length
	//如果handler未结束且Write的长度超过了最大写缓存量，我们使用chunk编码传输数据
	//会在finishRequest中，调用Flush之前将其设置成true
	handlerDone bool

	//bufw = bufio.NewBufioWriter(chunkWriter)
	bufw *bufio.Writer
	cw   *chunkWriter

	req *Request

	//是否在本次http请求结束后关闭tcp连接，以下情况需要关闭连接：
	//1、HTTP/1.1之前的版本协议
	//2、请求报文头部设置了Connection: close
	//3、在net.Conn进行Write的过程中发生错误
	closeAfterReply bool

	//是否使用chunk编码的方式，一旦检测到应该使用chunk编码，则会被chunkWriter设置成true
	chunking bool
}

// Write写入流的顺序：
// 用户在 handler 中对 ResponseWriter 写 => 对response写 => 对response的 bufw 成员写 =>
// bufw 是 chunkWriter 的封装，对 chunkWriter 写 => 对 (*chunkWriter).(*response).(*conn).bufw 写 =>
// 这个 bufw 是对 net.Conn 的封装，对 net.Conn 写。
func (w *Response) Write(p []byte) (int, error) {
	n, err := w.bufw.Write(p)
	if err != nil {
		w.closeAfterReply = true
	}
	return n, err
}

func (w *Response) Header() Header {
	return w.header
}

func (w *Response) WriteHeader(statusCode int) {
	if w.wroteHeader {
		return
	}
	w.statusCode = statusCode
	w.wroteHeader = true
}

type chunkWriter struct {
	resp  *Response
	wrote bool
}

func (cw *chunkWriter) Write(p []byte) (n int, err error) {
	//第一次触发Write方法
	if !cw.wrote {
		cw.finalizeHeader(p)
		if err = cw.writeHeader(); err != nil {
			return
		}
		cw.wrote = true
	}
	bufw := cw.resp.c.bufw
	//当Write数据超过缓存容量时，利用chunk编码传输，chunk编码格式见该系列(4)。
	if cw.resp.chunking {
		_,err = fmt.Fprintf(bufw,"%x\r\n",len(p))
		if err!=nil{
			return
		}
	}
	n,err = bufw.Write(p)
	if err == nil && cw.resp.chunking{
		_,err=bufw.WriteString("\r\n")
	}
	return n,err
}

//设置响应头部
func (cw *chunkWriter) finalizeHeader(p []byte) {
	header := cw.resp.header
	//如果用户未指定Content-Type，我们使用嗅探。因为嗅探算法并非重点，我们这里直接使用标准库提供的api
	if header.Get("Content-Type") == "" {
		header.Set("Content-Type", http.DetectContentType(p))
	}
	//如果用户未指定任何编码方式
	if header.Get("Content-Length") == "" && header.Get("Transfer-Encoding") == "" {
		// 因为Flush触发该Write
		if cw.resp.handlerDone {
			buffered := cw.resp.bufw.Buffered()
			header.Set("Content-Length", strconv.Itoa(buffered))
		} else {
			//因为超出缓存触发该Write
			cw.resp.chunking = true
			header.Set("Transfer-Encoding", "chunked")
		}
		return
	}
	if header.Get("Transfer-Encoding") == "chunked"{
		cw.resp.chunking = true
	}
}

//将响应头部发送
func (cw *chunkWriter) writeHeader() (err error) {
	codeString := strconv.Itoa(cw.resp.statusCode)
	//statusText是个map，key为状态码，value为描述信息，见status.go，拷贝于标准库
	statusLine := cw.resp.req.Proto + " " + codeString + " " + statusText[cw.resp.statusCode] + "\r\n"
	bufw := cw.resp.c.bufw
	_, err = bufw.WriteString(statusLine)
	if err != nil {
		return
	}
	for key, value := range cw.resp.header {
		_, err = bufw.WriteString(key + ": " + value[0] + "\r\n")
		if err != nil {
			return
		}
	}
	_, err = bufw.WriteString("\r\n")
	return
}