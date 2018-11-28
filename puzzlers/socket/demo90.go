package main

import (
	"fmt"
	"net"
	"syscall"
)

func main() {
	fd1, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		fmt.Printf("socket error: %v\n", err)
		return
	}
	defer syscall.Close(fd1)
	fmt.Printf("The file descriptor of socket: %d\n", fd1)

	conn, err := net.Dial("tcp", "1.1.1.1:80")
	_ = conn
}
