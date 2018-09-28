package main

import (
	"errors"
	"fmt"
)

type operate func(x, y int) int

// 1
func calculate(x, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

// 2
type calculateFunc func(x int, y int) (int, error)

func genCalculator(op operate) calculateFunc {
	return func(x int, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func main() {
	// 1
	x, y := 12, 23
	op := func(y, x int) int {
		return x + y
	}
	result, err := calculate(x, y, op)
	fmt.Printf("The result: %d (error: %v)\n", result, err)
	result, err = calculate(x, y, nil)
	fmt.Printf("The result: %d (error: %v)\n", result, err)

	// 2
	x, y = 56, 78
	add := genCalculator(op)
	result, err = add(x, y)
	fmt.Printf("The result: %d (error: %v)\n", result, err)
}
