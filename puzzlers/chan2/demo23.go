package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 1
	var uselessChan = make(chan<- int, 1)
	anotherUselessChan := make(<-chan int, 1)
	fmt.Printf("The useless channels: %v, %v\n", uselessChan, anotherUselessChan)

	// 2
	intChan1 := make(chan int, 3)
	SendInt(intChan1)

	// 4
	intChan2 := getIntChan()
	for elem := range intChan2 {
		fmt.Printf("The element in intChan2: %v\n", elem)
	}

	// 5
	_ = GetIntChan(getIntChan)
}

// 2
func SendInt(ch chan<- int) {
	ch <- rand.Intn(1000)
}

// 3
type Notifier interface {
	SendInt(ch chan<- int)
}

// 4
func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

// 5
type GetIntChan func() <-chan int
