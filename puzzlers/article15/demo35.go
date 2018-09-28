package main

type Named interface {
	Name() string
}

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func main() {
	// 1
	const num = 123
	//_ := &num // 常量不可寻址
	//_ := &(123) // 基本类型字面量不可寻址

	var str = "abc"
	_ = str
	//_ := &(str[0]) // 字符串变量的索引结果不可寻址
	//_ := &(str[0:2]) // 字符串变量的切片结果不可寻址
	str2 := str[0:2]
	_ = &str2

	//_ = &(123 + 456) // 算数表达式的结果不可寻址
	num2 := 456
	_ = num2
	//_ = &(num + num2) // 算数表达式的结果不可寻址

	//_ = &([3]int{1, 2, 3}[0]) // 对数组字面量的索引结果值不可寻址
	//_ = &([3]int{1, 2, 3}[0:2]) // 对数组字面量的切片结果值不可寻址
	_ = &([]int{1, 2, 3}[0])
	//_= &([]int{1, 2, 3}[0:2]) // 对切片字面量的切片结果不可寻址,runtime error
	//_ = &(map[int]string{1: "a"}[0]) // 对字典字面量的索引结果值不可寻址

	var map1 = map[int]string{1: "a", 2: "b", 3: "c"}
	_ = map1
	//_ = &(map1[2]) // 对字典变量的索引结果值不可寻址

	//_ = &(func(x, y int) int {
	//	return x + y
	//}) // 字面量代表的函数不可寻址
	//_ = &(fmt.Sprintf) // 标识符代表的函数不可寻址
	//_ = &(fmt.Sprintln("abc")) // 函数的调用结果不可寻址

	dog := Dog{name: "little pig"}
	_ = dog
	//_ = &(dog.Name) // 标识符代表的函数不可寻址
	//_ = &(dog.Name()) // 方法的调用结果值不可寻址

	//_ = &(Dog{"little pig"}.name) // 结构体字面量的字段不可寻址

	//_ = &(interface{}(dog)) // 类型转换的结果值不可寻址
	dogI := interface{}(dog)
	_ = dogI
	//_ = &(dogI.(Named)) // 类型断言的结果不可寻址
	named := dogI.(Named)
	_ = named
	//_ = &(named.(Dog)) // 类型断言的结果不可寻址

	chan1 := make(chan int, 1)
	chan1 <- 1
	//_ = &(<-chan1) // 接受表达式的结果不可寻址
}
