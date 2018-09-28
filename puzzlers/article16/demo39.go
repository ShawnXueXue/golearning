package main

import (
	"fmt"
)

func main() {
	num := 10
	sign := make(chan struct{}, num)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
			sign <- struct{}{}
		}()
	}

	// 1
	//time.Sleep(time.Millisecond * 500)

	// 2
	for j := 0; j < num; j++ {
		<-sign
	}
}
