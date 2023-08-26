package httpd

func setupResponse(c *conn) *Response {
	return &Response{c}
}

type ResponseWriter interface {
	Write([]byte) (n int, err error)
}

type Response struct {
	c *conn
}

func (r *Response) Write(p []byte) (n int, err error) {
	return r.c.bufw.Write(p)
}
