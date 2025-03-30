package main

import (
	"errors"
	"fmt"
)

func main() {
	testDefer()
}

// defer从触发点开始，由下往上执行，所以输出结果：
//
//	Error from panic
//	Error from defer
func testDefer() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		panic("Error from defer")
	}()

	panic(errors.New("Error from panic"))
}
