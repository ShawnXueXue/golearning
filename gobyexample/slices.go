package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func main() {
	s := make([]string, 3)
	fmt.Println("type:", reflect.TypeOf(s).String())
	fmt.Println("emp:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(index(s, "1"))
	fmt.Println(strings.Index("abc", "b"))
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Printf("  s len:%d,   s cap:%d\n", len(s), cap(s))
	s = append(s, "d")
	fmt.Printf("  s len:%d,   s cap:%d\n", len(s), cap(s))
	s = append(s, "e", "f")
	fmt.Printf("  s len:%d,   s cap:%d\n", len(s), cap(s))
	s = append(s, "g")
	fmt.Printf("  s len:%d,   s cap:%d\n", len(s), cap(s))
	fmt.Println("apd:", s)
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[1:5]
	fmt.Println("sl1:", l)
	l = s[0:5]
	fmt.Println("sl2:", l)
	l = s[2:]
	fmt.Println("sl3:", l)
	l = s[:]
	slice1 := (*Slice)(unsafe.Pointer(&l))
	slice2 := (*Slice)(unsafe.Pointer(&s))
	fmt.Printf("  s len:%d,   s cap:%d,   s ptr:%p,   s data ptr:%p\n", len(s), cap(s), &s, slice2.Ptr)
	fmt.Printf("sl4 len:%d, sl4 cap:%d, sl4 ptr:%p, sl4 data ptr:%p\n", len(l), cap(l), &l, slice1.Ptr)

	t := []string{"g", "h", "l"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)
}

type Slice struct {
	Ptr unsafe.Pointer
	len int
	cap int
}

func index(list []string, target string) int {
	for i, v := range list {
		if v == target {
			return i
		}
	}
	return -1
}
