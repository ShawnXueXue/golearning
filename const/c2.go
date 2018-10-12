package consttest

import (
	"fmt"
	"reflect"
	"unsafe"
)

const a = "1"
const leafPageElementSize = int(unsafe.Sizeof(leafPageElement{}))

func t1() {
	fmt.Println(pageHeaderSize)
	fmt.Println(leafPageElementSize)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(b)

	const (
		a = iota    //0
		b           //1
		c           //2
		d = "hello" // 3
		e           // 4
		f = 100     // 5
		g           // 6
		h = iota    // 7
		i           // 8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

	const (
		//左移0位
		j = 1 << iota
		//左移1位
		k = 8 << iota
		//左移2位
		l
		//左移3位
		m
	)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
	fmt.Println("m=", m)
}
