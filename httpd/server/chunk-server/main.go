package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/report", func(writer http.ResponseWriter, request *http.Request) {
		for key, vals := range request.Header {
			fmt.Println("key: ", key, ", values: ", vals)
		}
		fmt.Println("Transfer-Encoding: ", request.TransferEncoding)


		bodyData := make([]byte, 0)
		buf := make([]byte, 1024)
		for {
			n, err := request.Body.Read(buf)
			if err != nil {
				if errors.Is(err, io.EOF) {
					bodyData = append(bodyData, buf[:n]...)
					break
				}
				fmt.Println("read data error: ", err)
				break
			}
			if n <= 0 {
				fmt.Println("nothing to read")
				time.Sleep(time.Second)
			} else {
				bodyData = append(bodyData, buf[:n]...)
			}

		}

		//bodyData, _ = io.ReadAll(request.Body)

		fmt.Println("body: ", string(bodyData))

	})


	err := http.ListenAndServe(":8099", nil)
	if err != nil {
		fmt.Println("start failed: ", err)
	}
}
