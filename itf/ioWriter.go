package itf

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (n int, err error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func t1() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
