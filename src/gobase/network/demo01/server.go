package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		fmt.Printf("wait info from client %v\n", conn.RemoteAddr().String())

		n, err := conn.Read(buf)

		if err == io.EOF {
			fmt.Println("client has exited.")
		}

		if err != nil {
			fmt.Println("read from client error is ", err)
			return
		}

		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("start listenning server ...")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")

	if err != nil {
		fmt.Printf("listen err is %v", err)
		return
	}

	defer listen.Close()

	for {
		fmt.Println("wait client connectting ...")
		conn, err := listen.Accept()

		if err != nil {
			fmt.Printf("Accept err is %v", err)
		} else {
			fmt.Printf("Accept connection is successfully! And the conn is %v\nand client ip is %v", conn, conn.RemoteAddr().String())
		}

		go process(conn)
	}

	fmt.Printf("listen is %v", listen)
}
