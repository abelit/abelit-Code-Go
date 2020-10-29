package main

import (
	"net/http"
)

type HelloHandler struct{}

func (m *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Wolrd!"))
}

type AboutHandler struct {
	Name string
}

func (h *AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.Name))
}

// 将使用HandleFunc来处理
func WelcomeServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("welcome page"))
}

func main() {

	mh := HelloHandler{}

	ah := AboutHandler{
		Name: "about page",
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}

	//　注册handler到
	http.Handle("/hello", &mh)

	http.Handle("/about", &ah)

	// 使用HandleFunc
	http.HandleFunc("/home", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("home page"))
	})

	http.HandleFunc("/welcome", WelcomeServeHTTP)

	// HandlerFunc + Handle
	http.Handle("/new", http.HandlerFunc(WelcomeServeHTTP))

	server.ListenAndServe()
	// 等价于
	// http.ListenAndServe("localhost:8080", nil)
}
