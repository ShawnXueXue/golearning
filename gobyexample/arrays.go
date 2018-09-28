package main

import (
	"fmt"
	"reflect"
)

func whatAmI(i interface{}) {
	fmt.Println(reflect.TypeOf(i).String())
}

func main() {
	var a [5]int
	fmt.Println("emp:", a)
	fmt.Println(reflect.TypeOf(a).String())
	// whatAmI(a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])
	fmt.Println("len: ", len(a))

	b := []int{1, 2, 3, 4, 5}
	whatAmI(b)
	fmt.Println("dcl: ", b)
	b = append(b, 6)
	fmt.Println("apd: ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
