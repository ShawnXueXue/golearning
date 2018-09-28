package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

type t1 struct {
	a int
	b string
}

type t2 struct {
	a int
	b string
}

func (t t2) String() string {
	return fmt.Sprintf("<a:%v><b:%v>", t.a, t.b)
}

func main() {
	r := rect{3, 4}
	c := circle{5}
	measure(r)
	measure(c)

	tmp1 := t1{1, "1"}
	tmp2 := t2{2, "2"}
	fmt.Printf("t1:%v\n", tmp1)
	fmt.Printf("t2:%v", tmp2)
}
