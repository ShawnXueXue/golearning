package main

import "fmt"

func main() {
	str := "hello 你好"
	fmt.Println("len(str):", len(str))
	fmt.Println("rune:", len([]rune(str)))
}
