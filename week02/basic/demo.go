package main

import "fmt"

func main() {
	// testPointer()
	// testSlice()
	testDefer()
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

func testSlice() {
	arr := [...]int{1, 2, 3, 4, 5}
	// 生成的新的切片，截取操作符还可以有第三个参数
	// 形如 [i,j,k]，第三个参数 k 用来限制新切片的容量，但不能超过原数组（切片）的底层数组大小。
	// 截取获得的切片的长度和容量分别是：j-i、k-i。
	slice := arr[3:4:4]
	fmt.Printf("slice length = %d, slice capacity = %d, slice = %v", len(slice), cap(slice), slice)
}

func testDefer() {
	i := 5
	// 输出结果为5，因为这里会先保存一个i的副本，i是基础类型，是值传递。
	defer fmt.Println(i)
	i += 10
}

func testStructFunc() {
	p := Teacher{}
	// 输出结果：
	// 	showA
	// 	showB
	// 因为Teacher没有实现ShowA()，所以会递推调用内部的People的ShowA()
	p.ShowA()
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}
