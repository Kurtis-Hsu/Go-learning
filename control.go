package main

import "fmt"

func main() {
	a := 10

	if a > 20 {
		fmt.Println("if", "a > 20")
	} else if a > 10 {
		fmt.Println("if", "a > 10")
	} else {
		fmt.Println("if", "a <= 10")
	}

	switch a {
	case 10:
		fmt.Println("switch", a)
		fallthrough // go 语言需要手动开启 switch 贯穿，且只会贯穿一个 case
	case 20:
		fmt.Println("switch", a)
		break
	default:
		fmt.Println("switch", "default")
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		if i == 9 {
			break
		}
		fmt.Println("for", i)
	}
}
