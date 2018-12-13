package main

import "fmt"

func init() {
	fmt.Println("init func.")
}

func main() {
	fmt.Println("main func.")
	fmt.Println(add(1, 2))
	fmt.Println(add(1, 2, 3))
	fmt.Println(add([]int{1, 2, 3}...))
}

func add(args ...int) (sum int) {
	for _, arg := range args {
		sum += arg
	}
	return
}
