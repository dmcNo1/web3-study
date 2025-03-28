package main

import "fmt"

// 数组和切片
// 切片底层就是一个数组，切片中有一个指针指向这个数组，len表示切片的元素个数，cap表示底层数组的长度，原理类似Java的ArrayList
// 切片的本质，就是一个指向数组的指针
func main() {
	// testCreateSlice1()
	// testCreateSlice2()
	// testCreateSlice3()
	// testAppendSlice()
	// testCopy()
	// testError()
}

func testCreateSlice1() {
	arr := [5]int{1, 2, 3, 4, 5}
	// 从初始位置复制，到arr[1]的位置结束，底层数组就是arr[0]~arr结束的位置
	slice := arr[:2]
	fmt.Println("length slice  = ", len(slice))     // 2
	fmt.Println("capacity slice  = ", cap(slice))   // 5
	fmt.Printf("address of arr[0] = %p\n", &arr[0]) // 0xc00000e390
	fmt.Printf("address of arr = %p\n", &arr)       // 0xc00000e390
	fmt.Printf("address of slice = %p\n", slice)    // 0xc00000e390

	// 从arr[1]复制，到arr[3]的位置结束，底层数组就是arr[1]~arr结束的位置
	slice = arr[1:4]
	fmt.Println("length slice  = ", len(slice))     // 3
	fmt.Println("capacity slice  = ", cap(slice))   // 4
	fmt.Printf("address of arr[0] = %p\n", &arr[0]) // 0xc00000e390
	fmt.Printf("address of arr = %p\n", &arr)       // 0xc00000e390
	// 复制的位置不同，所以地址也不一样，会有一定的偏移
	fmt.Printf("address of slice = %p\n", slice) // 0xc00000e398
}

func testCreateSlice2() {
	// 源码：func make(t Type, size ...IntegerType) Type
	// t: 切片类型
	// size: 如果只传一个的话，就是len，cap默认和len相同；如果传两个的话，就是len，cap
	// 也可以这样理解：func make(t Type, len int, cap int)
	slice := make([]int, 3, 5)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func testCreateSlice3() {
	// 这种方式的话，默认len和cap相同，相当于slice := make([]int, 3, 3)
	slice := []int{1, 2, 3}
	fmt.Println(slice)
	fmt.Println(len(slice)) // 3
	fmt.Println(cap(slice)) // 3
}

func testAppendSlice() {
	// 触发扩容，会改变底层数组引用
	slice := []int{1, 2, 3}
	fmt.Printf("address of slice = %p\n", slice) // 0xc00001a120

	slice = append(slice, 4)
	fmt.Printf("address of slice = %p\n", slice) // 0xc00000e3c0

	// 没有触发扩容，底层数组引用不变
	slice = make([]int, 3, 5)
	fmt.Printf("address of slice = %p\n", slice) // 0xc00000e3f0

	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice = append(slice, 4)
	fmt.Printf("address of slice = %p\n", slice) // 0xc00000e3f0
}

func testCopy() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 5)
	copy(slice2, slice1)
	// slice2 = [1 2 3 0 0], length slice2 = 5, capacity slice2 = 5
	fmt.Printf("slice2 = %v, length slice2 = %v, capacity slice2 = %v\n", slice2, len(slice2), cap(slice2))

	slice2 = make([]int, 2)
	copy(slice2, slice1)
	// slice2 = [1 2], length slice2 = 2, capacity slice2 = 2
	// slice2的容量不够，所以只会复制前两个；换而言之，copy不会触发扩容机制
	fmt.Printf("slice2 = %v, length slice2 = %v, capacity slice2 = %v\n", slice2, len(slice2), cap(slice2))

	// 字符串底层是[]byte数组，所以也能支持切片操作
	str := "abcdefg"
	sliceStr1 := str[3:]
	fmt.Println(sliceStr1)
}

func testError() {
	// 只声明，没有创建的slice无法使用，这里会数组越界
	// var slice []int
	// slice[0] = 1
	// fmt.Println(slice)

	// slice无法相互比较，只能和nil比较，这样写会编译报错
	// slice1 := []int{1, 2, 3}
	// slice2 := []int{4, 5, 6}
	// fmt.Println(slice1 == slice2)
}
