package main

import (
	"fmt"
)

func main() {
	var names = make([]int, 5, 10)
	fmt.Println(names)
	fmt.Printf("The length and capablity of names slice is %v and %v\n", len(names), cap(names))

	names = append(names, 9, 10, 11)
	fmt.Println(names)

	fmt.Printf("The length and capablity of names slice is %v and %v\n", len(names), cap(names))

	names2 := names

	fmt.Println(names2)

	fmt.Printf("The memory address of names and names2 is %p and %p", names, names2)
}
