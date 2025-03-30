package main

import (
	"fmt"
)

func main() {
	str := "Q天音波"
	byteArr := []byte(str)
	fmt.Println("byteArr length =", len(byteArr)) // 10，GoLang中，默认使用UTF-8，一个汉字占用3个字节
	for _, b := range byteArr {
		fmt.Printf("%c", b)
	}
	fmt.Println()

	arr := []rune(str)
	fmt.Println("arr length =", len(arr)) // 4
	for _, c := range arr {
		fmt.Printf("%c", c)
	}
	fmt.Println()
}
