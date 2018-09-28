package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	fmt.Println("Enter function main.")

	defer func() {
		fmt.Println("Enter defer function.")
		// recover的正确用法
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
		fmt.Println("Exit defer function.")
	}()

	// recover的错误用法
	fmt.Printf("no panic: %v\n", recover())

	// panic
	panic(errors.New("something wrong"))

	// recover的错误用法
	p := recover()
	fmt.Printf("panic: %s\n", p)
	fmt.Println("Exit function main.")
}
