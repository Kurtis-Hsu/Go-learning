package main

import "fmt"

func main() {
	a := 10
	b := 3

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) // 3
	fmt.Println(a % b)
	a++ // a = a + 1
	fmt.Println(a)
	a = 100
	a-- // a = a - 1
	fmt.Println(a)

	a++
	a--

	// --a // 不支持前置

	c := 11
	d := 10
	fmt.Println(c == d)
	fmt.Println(c != d)
	fmt.Println(c > d)
	fmt.Println(c < d)
	fmt.Println(c >= d)
	fmt.Println(c <= d)

	t := true
	f := false

	fmt.Println(t && f)
	fmt.Println(t || f)
	fmt.Println(!t)

	a = 0b00111100
	b = 0b00001101

	// fmt.Println(t | f) // 布尔值不可以使用位运算
	fmt.Printf("%b\n", a&b)
	fmt.Printf("%b\n", a|b)
	fmt.Printf("%b\n", a^b)
	fmt.Printf("%b\n", a&^b)

	fmt.Printf("%b\n", a>>2)
	fmt.Printf("%b\n", b<<2)

	a = 1
	fmt.Println(a)
	a += 20
	a -= 15
	a *= 5
	a /= 10
	a %= 2
	a >>= 3
	a <<= 1
	a &= 0b11
	a |= 0b1111
	a ^= 0b1010

	p := &a // 取地址
	fmt.Println(p)
	fmt.Println(*p) // 指针指向的变量
}
