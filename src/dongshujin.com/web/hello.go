package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayHello(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm() // 解析参数，默认不会解析
	// 打印参数到后台输出
	fmt.Println(request.Form)
	fmt.Println("Scheme: ", request.URL.Scheme)
	fmt.Println("path: ", request.URL.Path)
	fmt.Println(request.Form["url_long"])

	for key, value := range request.Form {
		fmt.Println("key = ", key)
		fmt.Println("value = ", strings.Join(value, " "))
	}

	fmt.Fprintf(writer, "Hello Bennett")
}

func login(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Method: ", request.Method)
	if "GET" == request.Method {
		t, _ := template.ParseFiles("web/login.gtpl")
		t.Execute(writer, nil)
	} else {
		request.ParseForm()
		fmt.Println("Username: ", request.Form["username"])
		fmt.Println("Password: ", request.Form["password"])
	}
}

func RegisterHandler() {
	// 设置路由访问
	http.HandleFunc("/say/hello", sayHello)
	http.HandleFunc("/login", login)
	// 设置监听端口
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
