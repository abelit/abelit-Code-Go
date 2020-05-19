package utils

import (
	"fmt"
)

func Add(a int,b int)  int {
	return a + b
}

func init() {
	fmt.Println("This is init method in add.go under utils pacakge!")
}