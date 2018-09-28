package main

import (
	"errors"
	"fmt"
)

func echo(request string) (response string, err error) {
	if request == "" {
		err = errors.New("empty content")
	}
	response = fmt.Sprintf("echo: %s", request)
	return
}

func main() {
	// 1
	for _, req := range []string{"", "hello!"} {
		fmt.Printf("request: %s\n", req)
		response, err := echo(req)
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}
		fmt.Printf("response: %s\n", response)
	}
	fmt.Println()

	// 2
	err1 := fmt.Errorf("invalid content: %s", "#$%")
	err2 := errors.New(fmt.Sprintf("invalid content: %s", "#$%"))
	if err1.Error() == err2.Error() {
		fmt.Println("The error messages in err1 and err2 are the same.")
	}
}
