package main

import "fmt"

func main() {
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
