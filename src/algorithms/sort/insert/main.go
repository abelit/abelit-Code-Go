package main

import "fmt"

//插入排序
func InsertionASort(s []int) {
	n := len(s)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			swap(s, j, j-1)
		}
		fmt.Println(s)
	}
}

func InsertionZSort(s []int) {
	n := len(s)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if s[j] > s[j-1] {
				swap(s, j, j-1)
			}
		}
	}
}

func swap(slice []int, i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func main() {
	s := []int{9, 0, 6, 5, 8, 2, 1, 7, 4, 3}
	fmt.Println(s)
	InsertionASort(s)
	fmt.Println(s)
	// InsertionZSort(s)
	// fmt.Println(s)
}
