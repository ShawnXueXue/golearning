package main

import "fmt"

func main() {
	{
		// alias use = to declare
		type MyString = string
		str := "BCD"
		myStr1 := MyString(str)
		myStr2 := MyString("A" + str)
		fmt.Printf("%T(%q) == %T(%q): %v\n",
			str, str, myStr1, myStr1, str == myStr1)
		fmt.Printf("%T(%q) > %T(%q): %v\n",
			str, str, myStr2, myStr2, str > myStr2)
		fmt.Printf("Type %T is the same as type %T.\n", str, myStr1)

		strs := []string{"E", "F", "G"}
		myStrs := []MyString(strs)
		fmt.Printf("A value of type []MyString: %T(%q)\n", myStrs, myStrs)
		fmt.Printf("Type %T is tha same as type %T\n", strs, myStrs)
	}
	{
		type MyString string
		str := "BCD"
		myStr1 := MyString(str)
		myStr2 := MyString("A" + str)
		_ = myStr2
		//fmt.Printf("%T(%q) == %T(%q): %v\n",
		//	str, str, myStr1, myStr1, str == myStr1) // compile error:Invalid operation: str == myStr2 (mismatched types string and MyString)
		//fmt.Printf("%T(%q) > %T(%q): %v\n",
		//	str, str, myStr2, myStr2, str > myStr2) // compile error:Invalid operation: str > myStr2 (mismatched types string and MyString)
		fmt.Printf("Type %T is different from type %T\n", myStr1, str)

		strs := []string{"E", "F", "G"}
		// myStrs的潜在类型是MyString
		var myStrs []MyString
		//myStrs = []MyString(strs) // compile error:Cannot convert expression of type []string to type []MyString
		//fmt.Printf("A value of type []MyString: %T(%q)\n",
		//	myStrs, myStrs)
		fmt.Printf("Type %T is different from type %T\n", myStrs, strs)
		fmt.Println()
	}
	{
		// 别名类型
		type MyString1 = string
		// 潜在类型
		type MyString2 string
		str := "BCD"
		myStr1 := MyString1(str)
		myStr2 := MyString2(str)
		// 潜在类型相同的不同类型之间可以进行类型转换
		myStr1 = MyString1(myStr2)
		myStr2 = MyString2(myStr1)

		myStr1 = str
		// 潜在类型相同的两个类型,值不能判等或比较,他们的变量之间也不能复制
		//myStr2 = str // Cannot use 'str' (type string) as type MyString2 in assignment
		//myStr1 = myStr2 // Cannot use 'myStr2' (type MyString2) as type MyString1 in assignment
		//myStr2 = myStr1 // Cannot use 'myStr1' (type MyString1) as type MyString2 in assignment
	}
}
