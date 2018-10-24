package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	propergateTest()
	//closeTest()
}

func propergateTest() {
	rootCtx, cancel := context.WithCancel(context.Background())
	p1Ctx, _ := context.WithCancel(rootCtx)
	p2Ctx, _ := context.WithCancel(rootCtx)
	p1cCtx, _ := context.WithCancel(p1Ctx)
	go func() {
		_, ok := <-p1Ctx.Done()
		if !ok {
			fmt.Println("p1 cancel.")
		}
		//select {
		//case <-p1Ctx.Done():
		//	fmt.Println("p1 cancel.")
		//}
	}()
	go func() {
		_, ok := <-p2Ctx.Done()
		if !ok {
			fmt.Println("p2 cancel.")
		}
		//select {
		//case <-p2Ctx.Done():
		//	fmt.Println("p2 cancel.")
		//}
	}()
	go func() {
		_, ok := <-p1cCtx.Done()
		if !ok {
			fmt.Println("p1c cancel.")
		}
		//select {
		//case <-p1cCtx.Done():
		//	fmt.Println("p1c cancel.")
		//}
	}()
	time.Sleep(time.Millisecond * 500)
	cancel()
	<-rootCtx.Done()
	time.Sleep(time.Millisecond * 500)
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
