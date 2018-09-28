package main

import (
	"fmt"
)

func main() {
	//c := make(chan int, 10)
	//for i := 0; i < 10; i++ {
	//	c <- i
	//	go func() {
	//		fmt.Println(<-c)
	//	}()
	//}

	// 输出不可预测,可能是空,可能是乱序的重复的数字
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	//time.Sleep(time.Hour)
}
