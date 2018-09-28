package main

import (
	"fmt"
)

func zeroValue(ival int) {
	ival = 0
}

func zeroPtr(pval *int) {
	*pval = 0
}

type a struct {
	b int
}

func change(m []*a) {
	for _, e := range m {
		e.b = 9
	}
}

func main() {
	i := 1
	fmt.Println("init value:", i)
	zeroValue(i)
	fmt.Println("zeroValue:", i)
	zeroPtr(&i)
	fmt.Println("zeroPtr:", i)
	fmt.Println("pointer:", &i)

	m := make([]*a, 0)
	m = append(m, &a{b: 0})
	m = append(m, &a{b: 1})
	m = append(m, &a{b: 2})
	change(m)
	for _, x := range m {
		fmt.Printf("%d\n", x.b)
	}
}
