package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int, 7)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	//v1 := m["k1"]
	//fmt.Println("v1:", v1)
	//fmt.Println("len:", len(m))
	//delete(m, "k2")
	//fmt.Println("map:", m)
	//_, prs := m["k1"]
	//fmt.Println("prs:", prs)
	//
	//n := map[string]int{"foo": 1, "bar": 2}
	//fmt.Println("map:", n)
	//
	//im := make(map[string]interface{})
	//var a = "1"
	//var b = 2
	//im["1"] = a
	//im["2"] = b
	//for k, v := range im {
	//	fmt.Printf("key:%s, value:%v, type:%T\n", k, v, v)
	//}
	//
	//pm := make(map[*string]interface{})
	//str1 := "123"
	//str2 := "123"
	//pm[&str1] = 123
	//pm[&str2] = "321"
	//for k, v := range pm {
	//	fmt.Printf("key: %s, value: %v, type: %T\n", *k, v, v)
	//}
}
