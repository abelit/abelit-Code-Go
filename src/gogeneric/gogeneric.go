package main

import (
	"fmt"
	"reflect"
)

func Index[T comparable](s []T,x T) int {
	for i,v := range s {
		if v == x {
			return i
		}
	}
	return -1
}


func Add1(x,y int) int {
	return x + y
}

type Addable interface { 
	int |
	float64 |
	string |
	int64 |
	float32 |
	int32 |
	int16 |
	int8 |
	uint |
	uint64 |
	uint32 |
	uint16 |
	uint8 
}
func Add2[T Addable](x,y T) T {
	fmt.Println(reflect.TypeOf(x))
	return x + y
}

// func Add3(x,y interface{}) string {
// 	return fmt.Sprintf("%v + %v = %v",x,y,x.(int)+y)
// }

func main() {
	// si := []int{1,2,3,4,5}

	// fmt.Println(Index(si,3))

	// ss := []string{"a","b","c","d","e"}

	// fmt.Println(Index(ss,"f"))


	// fmt.Println("eeee" + "ffff")

	num1 := Add1(1,2)
	fmt.Println(num1)

	num2 := Add2(1.5,2.55)
	fmt.Println(num2)

	fmt.Println(Add2("hello","world"))
}