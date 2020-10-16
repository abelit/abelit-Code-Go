package main

import (
	"fmt"
	"reflect"
)

func TestReflect01(b interface{}) {
	rValue := reflect.ValueOf(b)

	fmt.Printf("rValue kind=%v\n", rValue.Kind())

	rValue.Elem().SetInt(20)

}

func main() {
	var num int = 10

	TestReflect01(&num)

	fmt.Println("num=", num)
}
