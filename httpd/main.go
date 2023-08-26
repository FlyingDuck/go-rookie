package main

import "net/http"

func main() {
	srv := http.Server{
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

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello net/http"))
}
