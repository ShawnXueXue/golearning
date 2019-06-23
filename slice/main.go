package main

import (
	"fmt"
	"strings"
	"time"
	"unsafe"
)

type Slice struct {
	Ptr unsafe.Pointer
	len int
	cap int
}

func main() {
	//array := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s1 := array[:5]
	//s2 := append(s1, 10)
	//fmt.Printf("s1 ptr: %p, s2 ptr: %p\n", &s1, &s2)
	//slice1 := (*Slice)(unsafe.Pointer(&s1))
	//slice2 := (*Slice)(unsafe.Pointer(&s2))
	//fmt.Printf("slice1 ptr: %p, slice2 ptr: %p\n", slice1.Ptr, slice2.Ptr)
	//
	//s3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s4 := append(s3, 10)
	//fmt.Printf("s3 ptr: %p, s4 ptr: %p\n", &s3, &s4)
	//slice3 := (*Slice)(unsafe.Pointer(&s3))
	//slice4 := (*Slice)(unsafe.Pointer(&s4))
	//fmt.Printf("slice3 ptr: %p, slice4 ptr: %p\n", slice3.Ptr, slice4.Ptr)
	//ip := "172.2.2.2"
	//fmt.Println(strings.Index(ip, "172.") == -1)
	//fmt.Println(ip[:4])
	//fmt.Println(time.Now().Hour())
	//err := getNil()
	//fmt.Printf("aaa: %v", err)
	defer fmt.Printf("end sync hostGroup, time cost: %f", Timing(time.Now()))
	time.Sleep(2 * time.Second)
	sql := make([]string, 0, 2)
	sql = append(sql, "123")
	sql = append(sql, "456")
	fmt.Printf(":::::::::%v\n", sql)
	fmt.Println(strings.EqualFold("k8s", "K8x"))
}

func Timing(start time.Time) float64 {
	return time.Now().Sub(start).Seconds()
}

//func getNil() int {
//	return nil
//}
