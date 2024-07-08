package init1

import "fmt"

// init 在 main 之前执行
func init() {
	fmt.Println("init1.go")
}

func Do() {
	fmt.Println("init1 do")
}
