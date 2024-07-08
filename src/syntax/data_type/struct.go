package main

import (
	"fmt"
)

// 结构体
func main() {
	p := Person{name: "person"}
	fmt.Println(p.name)
	change1(p)
	fmt.Println(p.name)
	change2(&p)
	fmt.Println(p.name)

	tom := Person{"tom"}
	tom.sayHi()

	zhang3 := Chinese{Person{"zhang3"}}
	zhang3.sayHi()

	check(tom)

	v1 := '中'
	v2 := "A"

	fmt.Printf("数据：%d，类型：%T\n", v1, v1)
	fmt.Printf("数据：%c，类型：%T\n", v1, v1)
	fmt.Printf("数据：%s，类型：%T\n", v2, v2)

	// 字符串连接
	fmt.Println("Hel" + "lo")

	// 转义字符
	fmt.Println("Hello\tworld")

	str := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println("legnth:", len(str))

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}

	for i, v := range str {
		fmt.Printf("char[%d]:%c\n", i, v)
	}
}

// 声明一个 int 的别名
type Int int

// 声明一个结构体，并给予类型名称
type Person struct {
	name string
}

func change1(p Person) {
	p.name = "change1" // 传递的是副本，无法修改原参数的值
}

func change2(p *Person) {
	p.name = "change2" // 传递的是地址，可以修改原参数的值
}

// 绑定方法，注意需要加 * 保证传的是地址，否则所有操作都会操作在副本上
// 小写字母开头的成员是私有，大写开头是公开
func (p *Person) sayHi() {
	fmt.Println("Hi, I'm", p.name)
}

type Chinese struct {
	Person // 继承 Person
}

type all interface{} // 所有结构体默认实现了空接口，空接口有个别名：any

func check(arg any) {
	v, ok := arg.(Person)
	if ok {
		fmt.Println(v.name)
	}
	fmt.Println(v, ok) // 查看接口的具体类型
}
