package main

import (
	"fmt"
)

func josephus(n int, k int) int {
	if n == 1 {
		return 1
	}
	// fmt.Println("the out is ", (josephus(n-1, k)+k-1)%n+1)
	return (josephus(n-1, k)+k-1)%n + 1
}

func main() {
	n := 14
	k := 2
	fmt.Println("The chosen place is ", josephus(n, k))
}
