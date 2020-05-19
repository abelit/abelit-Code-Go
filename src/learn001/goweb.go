package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello World!</h1>")
}

// func main() {
// 	http.HandleFunc("/", hello)
// 	http.ListenAndServe(":8000", nil)
// 	fmt.Println("Go web is running on localhost:8000")
// }
