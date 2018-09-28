package main

import (
	"flag"
	"fmt"
	"os"
)

var filePath string

func init() {
	flag.StringVar(&filePath, "filePath", "d:/defer.txt", "file will created in the given path")
}

func main() {
	flag.Parse()
	f := createFile(filePath)
	defer closeFile(f)
	writeFile(f)
}

func createFile(name string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
