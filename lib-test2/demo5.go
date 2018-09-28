package main

import (
	"flag"
	. "golearning/lib-test2/lib"
)

var name string

// redeclared in this block
// var InnerName = "OutterName"

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	Hello(name)
	// will use InnerName declared in golearning/lib-test2/lib
	Hello(InnerName)
}
