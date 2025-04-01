package main

import "fmt"

func main() {
	// 数组是值传递
	arr := [3]int{0, 1, 2}
	testArray(arr)
	fmt.Println(arr)

	// slice、map、channel都是引用传递
	slice := []int{0, 1, 2}
	testSlice(slice)
	fmt.Println(slice)
}

func testArray(arr [3]int) {
	arr[1] = 11
}

func testSlice(slice []int) {
	slice[1] = 11
}
