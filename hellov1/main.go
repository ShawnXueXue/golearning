package main

import (
	"fmt"
	"strings"
	"time"
)

// SayHello public function should have comment
func SayHello(repeat int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(strings.Repeat("Hello World!", repeat))
}

func main() {
	//SayHello(1)
	//for i := 1; i < 5; i++ {
	//	go SayHello(i)
	//}
	//time.Sleep(time.Hour)
	fmt.Println(time.Now().Format("20060102150405"))
	var aa []string
	fmt.Println(len(aa))
}
