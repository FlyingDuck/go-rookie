package httpd

import "net"

type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}

// Server 仅负责WEB服务器的启动逻辑，http 协议的解析交给 conn 模块处理
type Server struct {
	Addr    string  //监听地址
	Handler Handler // 处理 HTTP 请求的回调函数
}

func (s *Server) ListenAndServe() error {
	l, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	for {
		rwc, err := l.Accept() // 得到的TCP连接rwc(ReadWriteCloser)
		if err != nil {
			continue
		}

		c := newConn(rwc, s)
		go c.serve() // 为每一个链接开启一个 goroutine
	}
}
