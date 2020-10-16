package main

import (
	"fmt"
	"reflect"
)

// TestReflect01 , 基本数据类型的reflect
func TestReflect01(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("reflect type is ", rType)

	// reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Println("reflect value is ", rValue)

	fmt.Printf("the true type of rValue is %T", rValue)

	newValue := rValue.Int() + 30

	fmt.Println("newValue is ", newValue)

	// 将rvalue转为interface{}
	iValue := rValue.Interface()

	// 将interface{}转为所需要的数据类型
	iNum := iValue.(int)
	fmt.Printf("iNum is %v, and type is %T", iNum, iNum)
}

// TestReflect02 ，结构体的reflect
func TestReflect02(b interface{}) {
	// reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("reflect type is ", rType)

	// reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Printf("reflect value is %v and type is %T, and rValue kind is %v, and rType kind is %v\n", rValue, rValue, rValue.Kind(), rType.Kind())

	// 将rvalue转为interface{}
	iValue := rValue.Interface()
	fmt.Printf("iValue is %v and type is %T\n", iValue, iValue)
	// 将interface{}通过断言转成需要的类型
	stu, ok := iValue.(Student)

	if ok {
		fmt.Printf("stu name is %v, and age is %v, and type is %T", stu.Name, stu.Age, stu)
	}
}

// Student is a struct
type Student struct {
	Name string
	Age  int
}

func main() {
	// num := 20
	// TestReflect01(num)

	stu := Student{
		Name: "abelit",
		Age:  28,
	}

	TestReflect02(stu)
}
