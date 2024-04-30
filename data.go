package main

import "fmt"

var global = "global"       // 全局变量
const CONSTANT = "constant" // 常量
const PI, STRING, BOOL = 3.14, "string", true
const (
	C1 = 1
	C2 = 2
)

func main() {
	var v1 string = "v1" // 使用 var 关键字定义变量，变量名后面跟类型
	var v2 = "v2"        // go 拥有类型推断，可以不显式定义类型

	// 定义多个变量
	var (
		v3 int
		v4 = 0
		v5 = 'c'
		v6 = true
	)

	v3 = 1

	fmt.Println(v1, v2, v3, v4, v5, v6)

	a := 1 // := 可以在定义变量的同时赋值（只能在函数内部使用）

	// v3 := 2 // 已经定义过的变量不可以使用 :=

	fmt.Println(a)

	// 变量交换
	b := 2
	fmt.Println("交换前：", a, b)
	a, b = b, a
	fmt.Println("交换后：", a, b)

	// 同时接收函数的多个返回值
	a, b = test()
	fmt.Println("接收多返回值", a, b)

	// 匿名变量
	c, _ := test()
	_, d := test()
	fmt.Println("匿名变量", c, d)

	var global = "test"          // 可以与全局变量重名
	fmt.Println("变量作用域", global) // 就近原则，优先使用局部变量

	const (
		i1 = iota
		i2 = iota
		i3 = iota
	)

	fmt.Println("递归常量 iota", i1, i2, i3)

	const (
		i4 = iota
		i5 = iota
		i6 = iota
	)
	fmt.Println("iota 在 const 关键字出现时被重置为 0", i4, i5, i6)

	const (
		i7  = iota   // 0
		i8           // 1
		i9           // 2
		i10 = "test" // test
		i11          // test
		i12 = 100    // 100
		i13          // 100
		i14 = iota   // 7
		i15          // 8
	)

	fmt.Println("iota", i7, i8, i9, i10, i11, i12, i13, i14, i15)
}

func test() (int, int) {
	return 100, 200
}
