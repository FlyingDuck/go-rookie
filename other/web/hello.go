package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
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
		theDate := time.Date(2016, time.November, 10, 23, 0, 0, 0, time.UTC)
		fmt.Printf("Date is : %s\n", theDate.Local())
		t, _ := template.ParseFiles("src/dongshujin.com/httpd/login.gtpl")
		t.Execute(writer, nil)
	} else {
		// Way NO.1
		// 1) 不需要显式的调用 requst.ParseForm(), 会自动的调用
		// 2) request.FormValue("param")的方式，只会返回同名参数中的第一个
		//username := request.FormValue("username")
		//password := request.FormValue("password")
		//fmt.Println("Username: ", username, " password: ", password)

		// Way NO.2
		request.ParseForm()
		fmt.Println("Username: ", request.Form["username"])
		fmt.Println("Password: ", request.Form["password"])

		if len(request.Form["username"][0]) == 0 {
			fmt.Println("Username is blank")
		}

		if matched, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$", request.Form["username"][0]); !matched {
			fmt.Println("Username Must be Hanzi")
		}

		if matched, _ := regexp.MatchString("^([\\w\\.\\_]{2,10})@(\\w{1.}).([a-z]{2,4})$", request.Form["email"][0]); !matched {
			fmt.Println("Email formate is wrong")
		}

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
