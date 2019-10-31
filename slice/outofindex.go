package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	// 两种初始化方式，容量不同会导致是否再分配内存
	target := []string{"a", "d", "t"}
	//target := make([]string, 0, 8)
	//target = append(target, "a", "d", "t")

	target = append(target, "b")
	fmt.Printf("slice header address: %p\n", &target)
	fmt.Printf("slice data address: %p\n", &target[3])
	fmt.Println("len:", len(target), ", cap:", cap(target))
	t1(target, []string{"c", "b", "e", "f"})
	fmt.Printf("slice header address: %p\n", &target)
	fmt.Println(target)
	fmt.Println(len(target), " ", cap(target))
	p := unsafe.Pointer(uintptr(unsafe.Pointer(&target[3])) + unsafe.Sizeof("")*3)
	fmt.Println(*(*string)(p))
}

func t1(target []string, input []string) {
	fmt.Println("enter t1")
	fmt.Printf("slice header address: %p\n", &target)
	fmt.Printf("slice data address: %p\n", &target[3])
	fmt.Println("len:", len(target), ", cap:", cap(target))
	for _, line := range input {
		exist := false
		for _, temp := range target {
			if strings.EqualFold(line, temp) {
				exist = true
				break
			}
		}
		if !exist {
			target = append(target, line)
		}
	}
	fmt.Println("len:", len(target), ", cap:", cap(target))
	fmt.Println("exit t1")
}
