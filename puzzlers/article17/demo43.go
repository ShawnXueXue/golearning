package main

import "fmt"

func main() {
	// 1
	//value3 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	//switch value3[4] { // Duplicate case
	//case 0, 1, 2:
	//	fmt.Println("0 or 1 or 2")
	//case 2, 3, 4:
	//	fmt.Println("2 or 3 or 4")
	//case 4, 5, 6:
	//	fmt.Println("4 or 5 or 6")
	//}

	// 2
	value5 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value5[4] {
	case value5[0], value5[1], value5[2]: // Duplicate case value5[2]
		fmt.Println("0 or 1 or 2")
	case value5[2], value5[3], value5[4]:
		fmt.Println("2 or 3 or 4")
	case value5[4], value5[5], value5[6]:
		fmt.Println("4 or 5 or26")
	}

	// 3
	value6 := interface{}(byte(127))
	e, ok := value6.(byte)
	//value6 := byte(127)
	//e, ok := interface{}(value6).(byte)
	fmt.Printf("%b\n", e)
	fmt.Println(ok)

	// byte是uint8的别名类型
	switch t := value6.(type) {
	case uint8, uint16:
		fmt.Println("uint8 or uint16")
	case byte:
		fmt.Printf("byte")
	default:
		fmt.Printf("unsupported type: %T", t)
	}
}
