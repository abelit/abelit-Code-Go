package utils

import (
	"fmt"
)

func Minus(a int, b int) int {
	return a - b
}

func init() {
	fmt.Println("This is init method in minus.go under utils package!")
}
