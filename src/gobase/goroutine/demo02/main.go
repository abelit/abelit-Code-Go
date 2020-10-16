package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	numMap = make(map[int]int, 20)
	lock   sync.Mutex
)

func GetNumber(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	numMap[n] = res
	lock.Unlock()
}

func main() {
	for i := 1; i <= 20; i++ {
		go GetNumber(i)
	}

	time.Sleep(5 * time.Second)

	lock.Lock()
	for i, v := range numMap {
		fmt.Printf("map[%d]=%d \n", i, v)
	}
	lock.Unlock()
}
