package main

import (
	"fmt"
	"github.com/FlyingDuck/go-rookie/httpd"
	"io"
)

func main() {
	srv := httpd.Server{
		Addr:    ":8826",
		Handler: new(MyHandler),
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

type MyHandler struct {
}

func (h *MyHandler) ServeHTTP(w httpd.ResponseWriter, r *httpd.Request) {
	//// 用户的头部信息保存到buff中
	//buff := &bytes.Buffer{}
	//// 测试Request的解析
	//_, _ = fmt.Fprintf(buff, "[query]name=%s\n", r.Query("name"))
	//_, _ = fmt.Fprintf(buff, "[query]token=%s\n", r.Query("token"))
	//_, _ = fmt.Fprintf(buff, "[cookie]foo1=%s\n", r.Cookie("foo1"))
	//_, _ = fmt.Fprintf(buff, "[cookie]foo2=%s\n", r.Cookie("foo2"))
	//_, _ = fmt.Fprintf(buff, "[Header]User-Agent=%s\n", r.Header.Get("User-Agent"))
	//_, _ = fmt.Fprintf(buff, "[Header]Proto=%s\n", r.Proto)
	//_, _ = fmt.Fprintf(buff, "[Header]Method=%s\n", r.Method)
	//_, _ = fmt.Fprintf(buff, "[Addr]Addr=%s\n", r.RemoteAddr)
	//_, _ = fmt.Fprintf(buff, "[Request]%+v\n", r)
	//
	////手动发送响应报文
	//io.WriteString(w, "HTTP/1.1 200 OK\r\n")
	//io.WriteString(w, fmt.Sprintf("Content-Length: %d\r\n", buff.Len()))
	//io.WriteString(w, "\r\n")
	//
	//w.Write(buff.Bytes())

	buf, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	const prefix = "your message:"
	io.WriteString(w, "HTTP/1.1 200 OK\r\n")
	io.WriteString(w, fmt.Sprintf("Content-Length: %d\r\n", len(buf)+len(prefix)))
	io.WriteString(w, "\r\n")
	io.WriteString(w, prefix)
	w.Write(buf)
}
