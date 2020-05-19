package main

import (
	"fmt"
	"os"
)

func main() {
	dir := "./test/info.log"
	f, err := os.Stat(dir)

	if err != nil {
		fmt.Println(err)
		// os.MkdirAll(dir, 0666)
		os.Create(dir)
		return
	}

	fmt.Println(f.IsDir())
}
