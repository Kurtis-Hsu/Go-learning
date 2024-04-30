package main

import "fmt"

func main() {
	fmt.Println(add(1, 1))

	fmt.Println("变量一", "变量二")
	fmt.Println(swap("变量一", "变量二"))

	a := "test1"
	fmt.Println("arg", a)
	arg(a)
	fmt.Println("arg", a)

	// 函数也是一种类型
	fmt.Printf("%T\n", do)
	fmt.Printf("%T\n", handle)
	fmt.Printf("%T\n", process)
	fmt.Printf("%T\n", swap)
	fmt.Printf("%T\n", arg)
	fmt.Printf("%T\n", sum)

	fn := do // 定义函数类型变量并赋值
	fn()     // 调用函数变量

	func() {
		fmt.Println("匿名函数")
	}() // 匿名函数必须再声明时调用

	r1 := func(a, b int) int {
		return a + b
	}(1, 1)

	fmt.Println("result:", r1)

	call(func() int { return 1 })

	defer func() {
		fmt.Println("延迟函数1")
	}()

	func() {
		defer func() {
			fmt.Println("局部延迟函数1")
		}()

		defer func() {
			fmt.Println("局部延迟函数2")
		}()

		fmt.Println("局部延迟函数之后的代码")
	}()

	defer func() {
		fmt.Println("延迟函数2")
	}()

	fmt.Println("延迟函数之后的代码")
}

func add(a, b int) int { return a + b }

// 无参无返回值函数
func do() {
	fmt.Println("do sth")
}

// 有一个参数的函数
func handle(arg string) {}

// 有两个参数的函数
// 有一个返回值的函数
func process(arg1, arg2 string) string {
	result := arg1 + arg2
	return result
}

func swap(a, b string) (string, string) {
	return b, a
}

func arg(a string) {
	a = a + "test2"
}

// 可变参数（只能有一个可变参数，必须放在参数列表最后）
func sum(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

// 回调函数
func call(callback func() int) {
	fmt.Println("callback:", callback())
}
