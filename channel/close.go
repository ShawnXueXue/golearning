package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 3)
	go func() {
		ch <- time.Now().String()
		ch <- time.Now().String()
		ch <- time.Now().String()
		fmt.Println("close ch")
		close(ch)
	}()
	time.Sleep(2 * time.Second)
	for {
		select {
		case msg, open := <-ch:
			if open {
				fmt.Println(msg)
			} else {
				fmt.Println("ch closed.")
				return
			}
		}
	}
}
