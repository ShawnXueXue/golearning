package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	coordinateWithContext()
}

func coordinateWithContext() {
	total := 12
	var num int32
	fmt.Printf("The number: %d [with context.Context]\n", num)
	ctx, cancel := context.WithCancel(context.Background())
	for i := 1; i <= total; i++ {
		go addNum(&num, i, func() {
			if atomic.LoadInt32(&num) == int32(total) {
				cancel()
			}
		})
	}
	<-ctx.Done()
	fmt.Println("end.")
}

func addNum(numP *int32, id int, deferFunc func()) {
	defer deferFunc()
	for i := 0; ; i++ {
		currNum := atomic.LoadInt32(numP)
		newNum := currNum + 1
		time.Sleep(time.Millisecond * 200)
		if atomic.CompareAndSwapInt32(numP, currNum, newNum) {
			fmt.Printf("The number: %d [%d-%d]\n", newNum, id, i)
			break
		}
	}
}
