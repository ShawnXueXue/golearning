package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//propergateTest()
	closeTest()
}

func propergateTest() {
	rootCtx, cancel := context.WithCancel(context.Background())
	p1Ctx, _ := context.WithCancel(rootCtx)
	p2Ctx, _ := context.WithCancel(rootCtx)
	p1cCtx, _ := context.WithCancel(p1Ctx)
	go func() {
		select {
		case <-p1Ctx.Done():
			fmt.Println("p1 cancel.")
		}
	}()
	go func() {
		select {
		case <-p2Ctx.Done():
			fmt.Println("p2 cancel.")
		}
	}()
	go func() {
		select {
		case <-p1cCtx.Done():
			fmt.Println("p1c cancel.")
		}
	}()
	cancel()
	//<-rootCtx.Done()
	time.Sleep(time.Second * 1)
}

func closeTest() {
	c := make(chan int)
	go func() {
		for {
			select {
			case i, ok := <-c:
				if ok {
					fmt.Printf("get: %v\n", i)
				} else {
					fmt.Println("end")
					return
				}
			}
		}
	}()
	c <- 1
	c <- 1
	close(c)
	time.Sleep(time.Second)
}
