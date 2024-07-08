package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3} // 带元素个数定义的是数组
	fmt.Println("函数外数组", arr)
	changeArray(arr)
	fmt.Println("函数外数组", arr)

	sli1 := []int{1, 2, 3} // 不带元素个数定义的是切片
	fmt.Println("函数外切片", sli1)
	changeSlice(sli1)
	fmt.Println("函数外切片", sli1)

	// map
	map1 := map[int]string{1: "1", 2: "2"}
	fmt.Println(map1)

	sli2 := make([]int, 10)             // 动态切片
	map2 := make(map[string]string, 10) // 动态 map
	fmt.Println(sli2)
	fmt.Println(map2)

	s := String("")
	s.isEmpty()
}

// 数组传值时传递的是深拷贝副本的引用
func changeArray(arr [3]int) {
	fmt.Println("函数内数组", arr)
	arr[0] = 100
	fmt.Println("函数内数组", arr)
}

// 切片传值时传递的是引用
func changeSlice(sli []int) {
	fmt.Println("函数内切片", sli)
	sli[0] = 100
	fmt.Println("函数内切片", sli)
}

type String string

func (str *String) isEmpty() bool {
	return len(*str) == 0
}
