package lib

import (
	"fmt"
)

var InnerName = "shawn"

// go install golearning/lib-test2/lib
// get .a: lib.a
func Hello(name string) {
	fmt.Printf("Hello %s!\n", name)
}
