package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func gId() int {
	buf := make([]byte, 64)
	n := runtime.Stack(buf, false)
	idFields := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))
	id, err := strconv.Atoi(idFields[0])
	if nil != err {
		panic(fmt.Sprintf("can not get gorountine id: %v", err))
	}
	return id
}

func main() {
	fmt.Println("main", gId())
	wg := sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i, gId())
		}(i)
	}
	wg.Wait()
}
