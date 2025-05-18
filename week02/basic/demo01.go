package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	testCmd()
	// testPointer()
	// testSwitch()
	// testGoto()
	// testArray()
}

func testCmd() {
	// 通过os包获取命令行参数
	args := os.Args
	for _, val := range args {
		fmt.Printf("%v\n", val)
	}

	// flag包获取命令行参数，参数通过k-v的形式传输
	// 示例：./xxx.exe -name=jackpot -age=24
	name := flag.String("name", "zhangsan", "please enter your name!")
	age := flag.Int("age", 24, "please enter your age!")
	// 解析参数，注意返回值类型都是指针
	flag.Parse()
	fmt.Printf("name = %v, age = %v\n", *name, *age)
}

func testPointer() {
	num1 := 666
	p := &num1
	fmt.Println(num1)
	fmt.Println(p)
	num1 = 777
	fmt.Println(num1)
	fmt.Println(p)
	*p = 999
	fmt.Println(num1)
	fmt.Println(p)

	var arrPtr *[3]int
	arr := [...]int{1, 2, 3}
	arrPtr = &arr
	fmt.Println(arrPtr)
}

func testSwitch() {
	crazy := 4

	switch num := 4; num {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fmt.Println("work day")
	// 也能用变量，甚至是函数的返回结果
	case crazy:
		fmt.Println("crazy")
	default:
		fmt.Println("bad day")
	}

	// 也可以用表达式，但是用表达式的时候不要指定比对的变量
	switch num := 6; {
	case num >= 5 && num <= 7:
		fmt.Println("weekend")
	default:
		fmt.Println("bad day")
	}
}

func testGoto() {
	num := 1
	// outer:
	if num <= 10 {
		fmt.Println(num)
		num++
		goto outer
	}
outer:
	fmt.Println("come here")
}

// GoLang中，比较两个数组是否相等，会先比对两个数组的长度，如果长度相同，再逐一比对每个元素，都相同说明两个数组相等
func testArray() {
	arr1 := [...]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	fmt.Println("arr1 == arr2, ", arr1 == arr2) // true

	// 在GoLang中，数组长度也是数组类型的一部分，所以不同长度的数组无法比较，会编译错误
	// arr3 := [...]int{2, 3}
	// fmt.Println("arr1 == arr2, ", arr1 == arr3)

	arr4 := [...]int{1, 3, 4}
	fmt.Println("arr1 == arr2, ", arr1 == arr4) // false
}
