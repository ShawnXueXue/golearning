package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

var p = fmt.Printf

func main() {
	pt := point{1, 2}
	p("%v\n", pt)                     // {1 2}
	p("%+v\n", pt)                    // {x:1 y:2}  %+v 的格式化输出内容将包括结构体的字段名
	p("%#v\n", pt)                    // main.point{x:1, y:2}  %#v 形式则输出这个值的 Go 语法表示
	p("%T\n", pt)                     // main.point  需要打印值的类型
	p("%t\n", true)                   // true  bool
	p("%d\n", 123)                    // 123  十进制数字
	p("%b\n", 14)                     // 1110  二进制数字
	p("%c\n", 33)                     // !  输出给定整数的对应字符
	p("%x\n", 456)                    // 1c8  十六进制
	p("%f\n", 78.9)                   // 78.900000  十进制浮点数
	p("%e\n", 123400000.0)            // 1.234000e+08
	p("%E\n", 123400000.0)            // 1.234000E+08
	p("%s\n", "\"string\"")           // "string"
	p("%q\n", "\"string\"")           // "\"string\""  像 Go 源代码中那样带有双引号的输出
	p("%x\n", "this is hex")          // 7468697320697320686578  输出使用 base-16 编码的字符串，每个字节使用 2 个字符表示
	p("%p\n", &pt)                    // 0xc04205a080  输出一个指针的值
	p("|%6d|%6d|\n", 12, 345)         // |    12|   345|
	p("|%6.2f|%6.2f|\n", 1.2, 3.45)   // |  1.20|  3.45|
	p("|%-6.2f|%-6.2f|\n", 1.2, 3.45) // |1.20  |3.45  |
	p("|%6s|%6s|\n", "foo", "b")      // |   foo|     b|
	p("|%-6s|%-6s|\n", "foo", "b")    // |foo   |b     |

	s := fmt.Sprintf("a %s", "string")         // 格式化并返回一个字符串而不带任何输出
	fmt.Println(s)                             // a string
	fmt.Fprintf(os.Stderr, "an %s\n", "error") // an error  格式化并输出到 io.Writers而不是 os.Stdout

	str := "aAzZ"
	for _, c := range str {
		fmt.Println(c >= 65 && c <= 90)
	}
}
