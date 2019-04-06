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
	defer_call()
	fmt.Println("333 Helloworld")
}

func defer_call() {
	defer func() {
		fmt.Println("11111")
	}()
	defer func() {
		fmt.Println("22222")
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from r : ", r)
		}
	}()

	defer func() {
		fmt.Println("33333")
	}()

	fmt.Println("111 Helloworld")

	panic("Panic 1!")

	panic("Panic 2!")

	fmt.Println("222 Helloworld")
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
