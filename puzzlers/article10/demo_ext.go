package main

import "fmt"

type msg struct {
	id int
	data string
}

//func main() {
//	ch1 := make(chan msg, 1)
//	m1 := msg {id: 1, data: "one"}
//	fmt.Printf("input msg: %v, pointer: %p\n", m1, &m1)
//	ch1 <- m1
//	m1.id = 2
//	m1.data = "two"
//	fmt.Printf("input msg: %v, pointer: %p\n", m1, &m1)
//	r1 := <- ch1
//	fmt.Printf("result msg: %v, pointer: %p\n", r1, &r1)
//}

func main() {
	ch1 := make(chan *msg, 1)
	m1 := &msg {id: 1, data: "one"}
	fmt.Printf("input msg: %v, pointer pointer: %p, pointer: %p\n", *m1, &m1, m1)
	ch1 <- m1
	m1.id = 2
	m1.data = "two"
	fmt.Printf("input msg: %v, pointer pointer: %p, pointer: %p\n", *m1, &m1, m1)
	r1 := <- ch1
	fmt.Printf("result msg: %v, pointer pointer: %p, pointer: %p\n", *r1, &r1, r1)
}
