package main

import "fmt"

//func main() {
//	nums := []int{2, 3, 4}
//	sum := 0
//	for _, num := range nums {
//		sum += num
//	}
//	fmt.Println("sum:", sum)
//	for i, _ := range nums {
//		fmt.Println("index:", i)
//	}
//
//	kvs := map[string]string{"a": "apple", "b": "banana"}
//	for k, v := range kvs {
//		fmt.Printf("%s -> %s\n", k, v)
//	}
//	for k := range kvs {
//		fmt.Println("Key:", k)
//	}
//	for i, c := range "go" {
//		fmt.Println(i, c)
//	}
//}

type student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for i, stu := range stus {
		fmt.Printf("%p, %p\n", &stu, &stus[i])
		m[stu.Name] = &stu
	}
	fmt.Println(m["zhou"])
	fmt.Println(m["li"])
	fmt.Println(m["wang"])
	fmt.Println(m)
}
