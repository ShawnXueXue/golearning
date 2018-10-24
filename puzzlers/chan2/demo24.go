package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	example1()
	example2()
}

func example1() {
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	index := getRandInt(3)
	fmt.Printf("The index: %d\n", index)
	intChannels[index] <- index
	select {
	case elem := <-intChannels[0]:
		fmt.Printf("The first candidate case is seleted, the element is %d.\n", elem)
	case elem := <-intChannels[1]:
		fmt.Printf("The second candidate case is seleted, the element is %d.\n", elem)
	case elem := <-intChannels[2]:
		fmt.Printf("The third candidate case is seleted, the element is %d.\n", elem)
	default:
		fmt.Println("No candidate case is selected!")
	}
}

func getRandInt(limit int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(limit)
}

func example2() {
	intChan := make(chan int, 1)
	time.AfterFunc(time.Second, func() {
		close(intChan)
	})
	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("The candidate case is closed.")
			break
		}
		fmt.Println("The candidate case is selected.")
	}
}
