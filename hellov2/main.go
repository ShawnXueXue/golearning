package main

import (
	"fmt"
)

// HelloSpeaker public interface
type HelloSpeaker interface {
	SayHello()
}

type chinese struct {
	Name string
}

type american struct {
	Name string
}

func (c *chinese) SayHello() {
	fmt.Printf("你好世界!我是%s\n", c.Name)
}

func (c *chinese) ChangeNamePointerReveiver(newName string) {
	c.Name = newName
}

func (c chinese) ChangeNameValueReveiver(newName string) {
	c.Name = newName
}

// Repeater function
func Repeater(speaker HelloSpeaker, repeat int) {
	for i := 0; i < repeat; i++ {
		speaker.SayHello()
	}
}

func main() {
	// var c = &chinese{
	// 	Name: "张三",
	// }

	var _ HelloSpeaker = &chinese{Name: "张三"}

	c := &chinese{
		Name: "张三",
	}
	fmt.Println(c)    //&{张三}
	fmt.Println(&c)   //0xc42008e018
	fmt.Println(*&c)  //&{张三}
	fmt.Println(**&c) //{张三}
	fmt.Println(*c)   //{张三}
	fmt.Println(&*c)  //&{张三}
	// Repeater(c, 1)
	// c.ChangeNameValueReveiver("李四")
	// Repeater(c, 1)
	// c.ChangeNamePointerReveiver("李四")
	// Repeater(c, 1)

	// c := chinese{
	// 	Name: "张三",
	// }
	// fmt.Println(c)   // {张三}
	// fmt.Println(&c)  // &{张三}
	// fmt.Println(*&c) // {张三}

}
