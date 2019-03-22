package main

import (
	"fmt"
)

func main() {
	a1 := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("a1: %v (len: %d, cap: %d)\n", a1, len(a1), cap(a1))
	s9 := a1[1:4]
	fmt.Printf("s9: %v (len: %d, cap: %d)\n", s9, len(s9), cap(s9))
	for i := 1; i <= 5; i++ {
		s9 = append(s9, i)
		fmt.Printf("s9: %v (len: %d, cap: %d)\n", s9, len(s9), cap(s9))
	}
	fmt.Printf("a1: %v (len: %d, cap: %d)\n", a1, len(a1), cap(a1))

	fmt.Println()
	s := []int{0, 1, 2, 3, 4, 5}
	ss := s[4:5]
	fmt.Println(s)
	fmt.Println(ss)
	ss = append(ss, 55)
	ss[0] = 44
	fmt.Println(s)
	fmt.Println(ss)
	ss = append(ss, 66)
	ss[0] = 444
	fmt.Println(s)
	fmt.Println(ss)

}
