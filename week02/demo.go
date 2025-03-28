package main

import "fmt"

func main() {
	testPointer()
}

func testPointer() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}
	slice[2] = 22
	for k, v := range m {
		fmt.Printf("%v -> %v, *v = %v\n", k, v, *v)
	}
}
