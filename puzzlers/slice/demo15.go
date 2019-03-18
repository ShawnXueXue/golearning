package main

import "fmt"

func main() {
	// 1
	s1 := make([]int, 5)
	fmt.Printf("Length of s1: %d\n", len(s1))
	fmt.Printf("Capacity of s1: %d\n", cap(s1))
	fmt.Printf("Value of s1: %d\n", s1)
	s2 := make([]int, 5, 7)
	fmt.Printf("Length of s2: %d\n", len(s2))
	fmt.Printf("Capacity of s2: %d\n", cap(s2))
	fmt.Printf("Value of s2: %d\n", s2)
	fmt.Println()
	// 2
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("Length of s4: %d\n", len(s4))
	fmt.Printf("Capacity of s4: %d\n", cap(s4))
	fmt.Printf("Value of s4: %d\n", s4)
	fmt.Println()
	// 3
	s5 := s4[:cap(s4)]
	fmt.Printf("Length of s5: %d\n", len(s5))
	fmt.Printf("Capacity of s5: %d\n", cap(s5))
	fmt.Printf("Value of s5: %d\n", s5)
}
