package main

import (
	"fmt"
)

func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	
	if leftIndex > rightIndex {
		fmt.Println("Cannot find value from array")
		return 
	}
        middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		BinaryFind(arr,leftIndex,middle-1,findVal)
	} else if (*arr)[middle] < findVal {
		BinaryFind(arr,middle+1,rightIndex,findVal)
	} else {
		fmt.Printf("Yeah, find the number is %v and index is %v \n",findVal,middle)
	}

}

func main() {
	fmt.Println("Hello Algorithms")
        arr := [6]int{3,35,67,89,100,500}

	fmt.Println("The sortable array is ", arr)

	//num := 67
	var num int

	fmt.Scanf("%d",&num)

	BinaryFind(&arr,0,len(arr)-1,num) 
}
