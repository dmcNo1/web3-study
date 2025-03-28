package main

import "fmt"

// 闭包
func main() {
	// 由于testFunc存在，所以num即使在函数调用完还是存在的
	testFunc := createUpdater()
	fmt.Println(testFunc())
	fmt.Println(testFunc())
	fmt.Println(testFunc())
	fmt.Println(testFunc())

	// 重新执行了一次，所以num重新从0开始，可以理解为不是同一个函数对象，所以两个闭包作用域不同
	testFunc = createUpdater()
	fmt.Println(testFunc())
	fmt.Println(testFunc())
	fmt.Println(testFunc())
	fmt.Println(testFunc())
}

// 这个函数返回一个返回值为int类型的函数
func createUpdater() func() int {
	num := 0
	return func() int {
		num++
		return num
	}
}
