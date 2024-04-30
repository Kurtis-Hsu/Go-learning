package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3} // 带元素个数定义的是数组
	fmt.Println("函数外数组", arr)
	changeArray(arr)
	fmt.Println("函数外数组", arr)

	sli := []int{1, 2, 3} // 不带元素个数定义的是切片
	fmt.Println("函数外切片", sli)
	changeSlice(sli)
	fmt.Println("函数外切片", sli)
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
