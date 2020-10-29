package main

import (
	"net/http"
)

type MyHandler struct{}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Wolrd!"))
}

func main() {

	mh := MyHandler{}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &mh,
	}
	server.ListenAndServe()
	// 等价于
	// http.ListenAndServe("localhost:8080", nil)
}
