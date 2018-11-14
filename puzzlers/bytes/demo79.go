package main

import (
	"bytes"
	"fmt"
)

func main() {
	var contents string
	buffer1 := bytes.NewBufferString(contents)
	fmt.Printf("The length of new buffer with contents %s: %d\n", contents, buffer1.Len())
	fmt.Printf("The capacity of new buffer with contents %s: %d\n", contents, buffer1.Cap())
	fmt.Println()
}
