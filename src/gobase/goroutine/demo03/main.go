package main

import (
	"fmt"
)

var (
	numMap  = make(map[int]int, 20)
	intChan chan int
)

func GetNumber(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	intChan <- res
}

func main() {
	intChan = make(chan int, 20)
	for i := 1; i <= 20; i++ {
		go GetNumber(i)
	}

	fmt.Printf("intChan is %v", len(intChan))
}
