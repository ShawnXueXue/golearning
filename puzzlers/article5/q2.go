package main

import (
	"fmt"
)

var container = []string{"zero", "one", "two"}

// var ppFree = "2"

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	fmt.Printf(container[1])
}
