package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func main() {

	chunkReport()
	//pingReport()
}

func chunkReport() {
	pr, rw := io.Pipe()
	defer pr.Close()

	go func(){
		defer rw.Close()
		for i := 0; i < 100; i++ {
			rw.Write([]byte(fmt.Sprintf("line:%d\r\n", i)))
			time.Sleep(100*time.Millisecond)
		}
		fmt.Println("write all data")
	}()
	http.Post("http://localhost:8099/report","text/pain", pr)


	//time.Sleep(10*time.Second)
}


func pingReport() {
	_, err := http.Post("http://localhost:8099/report", "text/plain", strings.NewReader("nothing"))
	if err != nil {
		fmt.Println("post error: ", err)
		return
	}
}
