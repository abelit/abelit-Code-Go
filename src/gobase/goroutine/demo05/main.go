package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "abelit"

	go func(msg string) {
		fmt.Println(msg)
	}(msg)

	time.Sleep(10 * time.Microsecond)
}
