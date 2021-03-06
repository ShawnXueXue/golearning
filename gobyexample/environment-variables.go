package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Printf("|%-35s|%-90s|\n", pair[0], pair[1])
		fmt.Println(strings.Repeat("-", 128))
	}
}
