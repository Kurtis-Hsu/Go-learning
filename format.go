package main

import "fmt"

func main() {
	v := 1
	fmt.Printf("变量 v：%d，类型：%T，内存地址：%p\n", v, v, &v) // 取地址符号 &，go 语言只允许取地址，不允许地址运算

	v = 2
	fmt.Printf("变量 v：%d，类型：%T，内存地址：%p\n", v, v, &v)
}
