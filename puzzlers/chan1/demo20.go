package main

import "fmt"

func main() {
	ch1 := make(chan int, 3)
	fmt.Printf("len: %d, cap: %d\n", len(ch1), cap(ch1))
	ch1 <- 2
	fmt.Printf("len: %d, cap: %d\n", len(ch1), cap(ch1))
	ch1 <- 1
	fmt.Printf("len: %d, cap: %d\n", len(ch1), cap(ch1))
	ch1 <- 3
	fmt.Printf("len: %d, cap: %d\n", len(ch1), cap(ch1))
	ch1 <- 3
	fmt.Printf("len: %d, cap: %d\n", len(ch1), cap(ch1))
	elem1 := <-ch1
	fmt.Printf("len: %d, cap: %d\n", len(ch1), cap(ch1))
	fmt.Printf("The first element received from channel ch1: %v\n", elem1)
}
