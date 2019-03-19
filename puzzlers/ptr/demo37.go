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
