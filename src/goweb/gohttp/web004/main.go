package main

import (
	"net/http"
)

func main() {
	// http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	// 	http.ServeFile(res, req, "wwwroot"+req.URL.Path)
	// })
	// http.ListenAndServe(":8080", nil)

	http.ListenAndServe(":8080", http.FileServer(http.Dir("wwwroot")))
}
