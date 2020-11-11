package main

import "fmt"

func main() {
	numbers := []int{6, 2, 7, 5, 8, 9}
	SelectASort(numbers)

	fmt.Println(numbers)
	SelectZSort(numbers)

	fmt.Println(numbers)
}

func SelectASort(values []int) {
	length := len(values)
	if length <= 1 {
		return
	}

	for i := 0; i < length; i++ {
		min := i // 初始的最小值位置从0开始，依次向右

		// 从i右侧的所有元素中找出当前最小值所在的下标
		for j := length - 1; j > i; j-- {
			if values[j] < values[min] {
				min = j
			}
		}
		//fmt.Printf("i:%d min:%d\n", i, min)

		// 把每次找出来的最小值与之前的最小值做交换
		values[i], values[min] = values[min], values[i]

		//fmt.Println(values)
	}
}

func SelectZSort(values []int) {
	length := len(values)
	if length <= 1 {
		return
	}
	for i := 0; i < length; i++ {
		max := i

		for j := length - 1; j > i; j-- {
			if values[j] > values[max] {
				max = j
			}
		}
		values[i], values[max] = values[max], values[i]
	}
}
