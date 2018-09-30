package main

import (
	"fmt"
	"unsafe"
)

type Slice struct {
	ptr unsafe.Pointer
	len int
	cap int
}

func main() {
	array := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := array[:5]
	s2 := append(s1, 10)
	fmt.Printf("s1 ptr: %p, s2 ptr: %p\n", &s1, &s2)
	slice1 := (*Slice)(unsafe.Pointer(&s1))
	slice2 := (*Slice)(unsafe.Pointer(&s2))
	fmt.Printf("slice1 ptr: %p, slice2 ptr: %p\n", slice1.ptr, slice2.ptr)

	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s4 := append(s3, 10)
	fmt.Printf("s3 ptr: %p, s4 ptr: %p\n", &s3, &s4)
	slice3 := (*Slice)(unsafe.Pointer(&s3))
	slice4 := (*Slice)(unsafe.Pointer(&s4))
	fmt.Printf("slice3 ptr: %p, slice4 ptr: %p\n", slice3.ptr, slice4.ptr)
}
