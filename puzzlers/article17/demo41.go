package main

import "fmt"

func main() {
	// 1
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	for i, _ := range numbers1 {
		if i == 3 {
			numbers1[i] |= i
		}
	}
	fmt.Println(numbers1)
	fmt.Println()

	// 2
	// [...]是数组的声明格式
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, v := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += v
		} else {
			numbers2[i+1] += v
		}
	}
	fmt.Println(numbers2)
	fmt.Println()

	// 3
	// range 操作的是变量的复制,切片是引用类型,改变原变量的同时,range变量也会改变
	numbers3 := []int{1, 2, 3, 4, 5, 6}
	maxIndex3 := len(numbers2) - 1
	for i, e := range numbers3 {
		if i == maxIndex3 {
			numbers3[0] += e
		} else {
			numbers3[i+1] += e
		}
	}
	fmt.Println(numbers3)
}
