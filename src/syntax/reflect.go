package main

import (
	"fmt"
	"reflect"
)

func main() {
	Value('a')

	user := User{1, "test", 18}
	Reflect(user)
}

func Value(arg any) {
	fmt.Println(reflect.TypeOf(arg))
	fmt.Println(reflect.ValueOf(arg))
}

type User struct {
	Id   int    `info:"id" desc:"ID"`
	Name string `info:"name" desc:"名字"`
	Age  int    `info:"age" desc:"年龄"`
}

func (u User) Print() {
	fmt.Println(u)
}

func Reflect(input any) {
	// 获取 input 的 type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType:", inputType)

	// 获取 input 的 value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue:", inputValue)

	// 通过 type 获取字段
	// 1. 获取 reflect.Type，通过 Type 得到 NumField 并遍历
	// 2. 得到每个 field 数据类型
	// 3. 通过 field Interface() 方法得到对应的 value

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 通过 type 获取方法

	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)

		fmt.Printf("%s: %v\n", method.Name, method.Type)
	}

	for i := 0; i < inputType.NumField(); i++ {
		info := inputType.Field(i).Tag.Get("info")
		desc := inputType.Field(i).Tag.Get("desc")
		fmt.Println("tag{ info:", info, ", desc:", desc, "}")
	}
}
