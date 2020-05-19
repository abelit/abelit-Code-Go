package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(rand.Int63n(10))
}
