package main

import (
	"fmt"
	"time"
)

func main() {
	errorMsg := make(chan string, 4)
	go func() {
		errorMsg <- "problem"
	}()
	time.Sleep(700 * time.Millisecond)
	select {
	case msg := <-errorMsg:
		fmt.Println("get error:" + msg)
	default:
		fmt.Println("ok")
	}
}
