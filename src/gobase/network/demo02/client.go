package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}

	fmt.Println("conn successfully and conn is ", conn)
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("read err is ", err)
		}

		if lineLabel := strings.Trim(line, " \r\n"); lineLabel == "exit" {
			fmt.Println("exit ...")
			break
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("content err is ", err)
		}
	}
	// fmt.Printf("client has sent %d bytes data.", content)
}
