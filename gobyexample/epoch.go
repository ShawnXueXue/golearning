package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
0000 0000-0000 007F | 0xxxxxxx
0000 0080-0000 07FF | 110xxxxx 10xxxxxx
0000 0800-0000 FFFF | 1110xxxx 10xxxxxx 10xxxxxx
0001 0000-0010 FFFF | 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
*/

func main() {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 1))
	fmt.Println(time.Unix(1, nanos))
	//cn := "𠑗"
	//cn := "𧺯"
	cn := "丁"
	bb := []byte(cn)
	for _, i := range bb {
		fmt.Printf("%x-", i)
	}
	fmt.Println()
	fmt.Println(strconv.QuoteToASCII(cn))
	fmt.Println(len(cn))
	fmt.Println(len(([]rune(cn))))
}
