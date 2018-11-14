package main

import (
	"bufio"
	"fmt"
	"strings"
)

var pf = fmt.Printf
var pl = fmt.Println

func main() {
	comment := "Package bufio implements buffered I/O. " +
		"It wraps an io.Reader or io.Writer object, " +
		"creating another object (Reader or Writer) that " +
		"also implements the interface but provides buffering and " +
		"some help for textual I/O."
	basicRaeder := strings.NewReader(comment)
	pf("The size of basic reader: %d\n", basicRaeder.Size())
	pl()

	pl("New a buffer reader...")
	reader1 := bufio.NewReader(basicRaeder)
	pf("The default size of buffered reader: %d\n", reader1.Size())
	pf("The number of unread bytes in the buffer: %d\n", reader1.Buffered())
	pl()

	bytes, err := reader1.Peek(7)
	if err != nil {
		pf("error: %v\n", err)
	}
	pf("Peeked contents(%d): %q\n", len(bytes), bytes)
	pf("The number of unread bytes in the buffer: %d\n", reader1.Buffered())
	pl()

	buf1 := make([]byte, 7)
	n, e := reader1.Read(buf1)
	if e != nil {
		pf("error: %v\n", e)
	}
	pf("Read contents(%d): %q\n", n, buf1)
	pf("The number of unread bytes in the buffer: %d\n", reader1.Buffered())
	pl()

	pf("Reset the basic reader (size: %d)...\n", len(comment))
	basicRaeder.Reset(comment)
	pf("Reset the buffered reader (size: %d)...\n", reader1.Size())
	reader1.Reset(basicRaeder)
	peekNum := len(comment) + 1
	pf("Peek %d bytes...\n", peekNum)
	bytes, err = reader1.Peek(peekNum)
	if err != nil {
		pf("error: %v\n", err)
	}
	pf("Peeked contents(%d): %q\n", len(bytes), bytes)
	pf("The number of peeked bytes: %d\n", len(bytes))
	pl()

	pf("Reset the basic reader (size: %d)...\n", len(comment))
	basicRaeder.Reset(comment)
	size := 300
	pf("New a buffered reader of size: %d...\n", size)
	reader2 := bufio.NewReaderSize(basicRaeder, size)
	peekNum = size + 1
	pf("Peek %d bytes...\n", peekNum)
	bytes, err = reader2.Peek(peekNum)
	if err != nil {
		pf("error: %v\n", err)
	}
	pf("Peeked contents(%d): %q\n", len(bytes), bytes)
	pf("The number of peeked bytes: %d\n", len(bytes))
}
