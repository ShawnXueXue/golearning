package main

import (
	"fmt"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done:", input)
}

//func main() {
//	min, max := 0, 100
//	fmt.Printf("%d--%d", min, max)
//	for min < max  {
//		i := (min + max) / 2
//		fmt.Printf("<=%d(y or n)  ", i)
//		var s string
//		//fmt.Scanf("%s", &s)
//		//fmt.Scan(&s)
//		fmt.Scanln(&s)
//		if s != "" && s[0] == 'y' {
//			max = i
//		} else {
//			min = i + 1
//		}
//	}
//	fmt.Printf("%d\n", max)
//}
