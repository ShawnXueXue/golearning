package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
		time.Sleep(time.Second * 3)
		c1 <- "three"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
		time.Sleep(time.Second * 4)
		c2 <- "four"
		time.Sleep(time.Second * 1)
		close(c2)
	}()
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("c1 received", msg1)
		case msg2, open := <-c2:
			if open {
				fmt.Println("c2 received", msg2)
			} else {
				fmt.Println("c2 closed, exit.")
				goto Exit
			}
		default:
			time.Sleep(time.Second * 1)
			fmt.Println("waiting...")
		}
	}
Exit:
	fmt.Println("exit select.")
}
