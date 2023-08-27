package httpd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type conn struct {
	svr *Server  // 引用上层服务处理器
	rwc net.Conn // 底层 tcp 链接

	// io.LimitedReader，它包含一个属性N代表能够在这个reader上读取的最多字节数，
	// 如果在此reader上读取的总字节数超过了上限， 则接下来对这个reader的读取都会
	// 返回 io.EOF，从而有效终止读取过程，避免首部字段的无限读。
	lr *io.LimitedReader
	// bufr 是一个 bufio.Reader，其底层的reader为上述的 io.LimitedReader。
	// 可以使用 ReadLine 方法方便的进行逐行读取。
	bufr *bufio.Reader
	// bufw 是一个 bufio.Writer，提供一个写入数据的缓冲区，避免写数据时频繁的
	// IO 操作
	bufw *bufio.Writer
}

func newConn(rwc net.Conn, svr *Server) *conn {
	lr := &io.LimitedReader{
		R: rwc,
		N: 1 << 20,
	} // 1MB
	return &conn{
		svr: svr,
		rwc: rwc,

		bufw: bufio.NewWriterSize(rwc, 4<<10), // 缓冲区大小 4KB
		lr:   lr,
		bufr: bufio.NewReaderSize(lr, 4<<10),
	}
}

func (c *conn) serve() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic recovered with err: %v", err)
		}
		c.close()
	}()

	// http1.1支持keep-alive长连接，所以一个连接中可能读出个请求，因此实用for循环读取
	req, err := c.readRequest()
	if err != nil {
		handleErr(err, c)
		return
	}

	resp := c.setupResponse()

	c.svr.Handler.ServeHTTP(resp, req)

	// 清空缓冲区
	if err := req.finish(); err != nil {
		return
	}


}

func (c *conn) readRequest() (*Request, error) {
	return readRequest(c)
}

func (c *conn) setupResponse() ResponseWriter {
	return setupResponse(c)
}

func (c *conn) close() {
	c.rwc.Close()
}

func handleErr(err error, c *conn) {
	fmt.Println(err)
}
