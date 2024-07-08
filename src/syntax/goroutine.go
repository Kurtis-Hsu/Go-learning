package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个协程，执行 newTask
	go newTask()

	i := 0

	for {
		i++
		fmt.Printf("main Goroutine %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func newTask() {
	i := 0

	for {
		i++
		fmt.Printf("new Goroutine %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
