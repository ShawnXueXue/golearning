package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(10)
	buf := make(chan string, 1)
	messages := make(chan string)
	buf <- "pong"
	messages <- "ping"
	select {
	case s := <-buf:
		fmt.Println(s)
	case ss := <-messages:
		fmt.Println(ss)
	}
}
