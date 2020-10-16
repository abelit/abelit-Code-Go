package main

import (
	"fmt"
	"time"
)

// GetFactorial is a function
func GetFactorial(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res = res * i
	}

	fmt.Printf("The factorial of %d is %d.\n", n, res)
}

// GetError is function
func GetError() {
	// recover使协程产生的错误不相互影响
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("GetError and err is ", err)
		}
	}()

	fmt.Println("GetError function is execution.")

	var myMap map[string]string

	myMap["name"] = "abelit"

}

func main() {
	go GetFactorial(10)
	go GetError()

	time.Sleep(3 * time.Second)

	for i := 0; i < 10; i++ {

		fmt.Println("main func is execution and the number is ", i)
		time.Sleep(10 * time.Microsecond)
	}
}
