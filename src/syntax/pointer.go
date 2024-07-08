package main

import "fmt"

func main() {
	v := 1
	fmt.Printf("变量 v：%d，类型：%T，内存地址：%p\n", v, v, &v) // 取地址符号 &，go 语言只允许取地址，不允许地址运算

	v = 2
	fmt.Printf("变量 v：%d，类型：%T，内存地址：%p\n", v, v, &v)

	a := 1
	b := 0
	fmt.Printf("变量 a：%d，类型：%T，内存地址：%p\n", a, a, &a)
	fmt.Printf("变量 b：%d，类型：%T，内存地址：%p\n", b, b, &b)

	fmt.Println("交换")
	swap2(&a, &b)
	fmt.Printf("变量 a：%d，类型：%T，内存地址：%p\n", a, a, &a)
	fmt.Printf("变量 b：%d，类型：%T，内存地址：%p\n", b, b, &b)
}

func swap2(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
