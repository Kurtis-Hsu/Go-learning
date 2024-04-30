package main

import "fmt"

func main() {
	// 基本类型
	var s string
	fmt.Printf("类型：%T, 默认值：%s\n", s, s)
	var b bool
	fmt.Printf("类型：%T, 默认值：%t\n", b, b)
	var ui uint
	fmt.Printf("类型：%T, 默认值：%d\n", ui, ui)
	var i int
	fmt.Printf("类型：%T, 默认值：%d\n", i, i)
	var f32 float32
	fmt.Printf("类型：%T, 默认值：%f\n", f32, f32)
	// 类型后面的数字表示占用字节数，float32 表示占用字节数为 32 float 类型，int64 表示占用字节数为 64 的 int 类型

	d1 := 1
	fmt.Printf("数据：%d，类型：%T\n", d1, d1)
	d2 := float64(d1) // 类型转换
	fmt.Printf("数据：%f，类型：%T\n", d2, d2)

	// bl := bool(d1) // 整型不能转换为布尔型
	v := string(d1) // 整型转换为字符串时会转换为码点而不是与数字相同的字符串
	fmt.Printf("数据：%s，类型：%T\n", v, v)
}
