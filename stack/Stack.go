package main

import "fmt"

type node struct {
	val  interface{}
	prev *node
	next *node
}
type Stack struct {
	first *node
	last  *node
}

func (s *Stack) Push(v interface{}) {
	n := &node{v, s.last, nil}
	if nil != s.last {
		s.last.next = n
	}
	s.last = n
}
func (s *Stack) Pop() interface{} {
	if nil == s.last {
		panic("Empty Stack.")
	}
	result := s.last.val
	if nil != s.last.prev {
		s.last.prev.next = nil
	}
	s.last = s.last.prev
	return result
}

type temp struct {
	a int
}

func main() {
	s := Stack{}
	s.Push(temp{})
	s.Push(temp{})
	s.Push(temp{})
	fmt.Println(s.Pop().(temp).a)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	m := make(map[string]interface{})
	if m["1"] == nil {
		fmt.Println("Y")
	} else {
		fmt.Println("N")
	}
	var aa []string
	aa = append(aa, aa...)
}
