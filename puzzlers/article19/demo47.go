package main

import "fmt"

func main() {
	//s1 := []int{0, 1, 2, 3, 4}
	//e5 := s1[5]
	//_ = e5

	for _, value := range []int{0, 1, 2, 3, 4} {
		value *= 2
		fmt.Println(value)
	}
}
