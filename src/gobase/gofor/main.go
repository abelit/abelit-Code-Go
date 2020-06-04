package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	count := 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		if n == 99 {
			break
		}
		count++

	}
	fmt.Printf("第%v次就发现了99\n", count)
}
