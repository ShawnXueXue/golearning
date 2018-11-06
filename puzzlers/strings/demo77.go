package main

import (
	"fmt"
	"strings"
)

func main() {
	reader1 := strings.NewReader("NewReader returns a new Reader reading from s. " +
		"It is similar to bytes.NewBufferString but more efficient and read-only.")
	fmt.Printf("The size of reader: %d\n", reader1.Size())
	fmt.Printf("The reading index in reader: %d\n", reader1.Size()-int64(reader1.Len()))
}
