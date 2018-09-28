package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	// flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
	// name := flag.String("name", "everyont", "The greeting object.")
}

func main() {
	// flag.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }
	flag.Parse()
	fmt.Printf("Hello %s!\n", name)
}
