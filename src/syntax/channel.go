package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		fmt.Println("run goroutine")

		c <- 666
	}()

	fmt.Println("number:", <-c)

	buf := make(chan int, 3) // 带缓存的通道

	go func() {
		fmt.Println("chan 1")
		c <- 1
	}()

	go func() {
		fmt.Println("chan 2")
		c <- 2
	}()

	go func() {
		fmt.Println("chan 3")
		c <- 2
	}()

	// go func() {
	// 	fmt.Println("chan 4")
	// 	c <- 2
	// }()

	time.Sleep(1000)
	fmt.Println("get chan", <-buf)
	fmt.Println("get chan", <-buf)
	fmt.Println("get chan", <-buf)
	// fmt.Println("get chan", <-buf)
}
