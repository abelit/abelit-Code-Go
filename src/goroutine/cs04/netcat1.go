package main

import (
	"io"
	"log"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法
func (r MyReader) Read(b []byte) (int, error) {
    // 赋值并返回
    b[0] = 'a'
	b[1] = 'b'
    return 1, nil
}

func main() {
	// conn, err := net.Dial("tcp","localhost:8000")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer conn.Close()
	
	// mustCopy(os.Stdout, strings.NewReader("Hello, World\n"))
	mustCopy(os.Stdout, MyReader{})
}