package main

import "fmt"

func main() {
	// 1
	// switch 和case的类型必须一致
	//value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	//switch 1 + 3 {
	//case value1[0], value1[1]:// Invalid case 'value1[1]' in switch on '1 + 3' (mismatched types 'int8' and 'int')
	//	fmt.Println("0 or 1")
	//}

	// 2
	// 如果case表达式中子表达式的结果值是无类型的常量,那么它的类型会被自动地转换为switch表达式的结果类型,又由于上述那几个整数都可以被转换为int8类型的值
	// 所以对这些表达式的结果值进行判等操作是没有问题的
	value2 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value2[1] {
	case 0, 1:
		fmt.Println("0 or 1")
		fallthrough // 紧挨在它下边的那个case子句附带的语句也会被执行
	case 2, 3:
		fmt.Println("2 or 3")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}
}
