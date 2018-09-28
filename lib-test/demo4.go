package main

import (
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

// go build golearning/lib-test
// get exe: lib-test.exe
func main() {
	flag.Parse()
	hello(name)
}
