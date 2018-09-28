package main

func main() {
	//var badMap1 = map[[]int]int{} // 会引发编译错误
	//_ = badMap1

	var badMap2 = map[interface{}]int{
		"1": 1,
		[]int{2}: 2, // 会引发panic
		3: 3,
	}
	_ = badMap2

	//var badMap3 map[[1][1]string]int // 会引发编译错误
	//_ = badMap3

	//type BadKey1 struct {
	//	slice []string
	//}
	//var badMap4 map[BadKey1]int // 会引发编译错误
	//_ = badMap4

	//var badMap5 map[[1][2][3][]string]int // 会引发编译错误
	//_ = badMap5

	//type BadKey2Field1 struct {
	//	slice []string
	//}
	//type BadKey2 struct {
	//	field BadKey2Field1
	//}
	//var badMap6 map[BadKey2]int // 会引发编译错误
	//_ = badMap6
}
