package main

import (
	"fmt"
)

func sum(a, b int) int {
	return a + b
}

func main() {
	a := 10
	b := 3

	res := sum(a, b)

	fmt.Println("Sum of ", a, " and ", b, " is ", res)
}
