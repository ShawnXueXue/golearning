package main

import (
	"fmt"
	"unsafe"
)

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func main() {
	// 1
	dog := Dog{"little pig"}
	dogP := &dog
	/**
	错误。在转换操作前，不能将uintptr的值保存到临时变量中。
	u := uintptr(p)
	p = unsafe.Pointer(u + offset)
	某些GC会把变量进行搬移来进行内存整理，这种类型的GC称为“移动的垃圾回收器”。当一个变量在内存中移动后，所有指向旧地址的指针都应该更新，并指向新地址。uintptr只是个数值，它的值不会变动。
	*/
	nameP := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(dogP)) + unsafe.Offsetof(dog.name)))
	fmt.Printf("nameP == &(dogP.name)? %v\n", nameP == &(dogP.name))
	fmt.Printf("The name of dog is %q.\n", *nameP)

	*nameP = "monster"
	fmt.Printf("The name of dog is %q.\n", dogP.name)
	fmt.Println()

	// 2
	numP := (*int)(unsafe.Pointer((uintptr(unsafe.Pointer(dogP)) + unsafe.Offsetof(dog.name))))
	num := *numP
	fmt.Printf("This is an unexpected number: %d\n", num)
}
