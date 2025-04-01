package main

import "fmt"

type Person interface {
	read()
}

type Student struct {
	name string
	age  int
}

func (s Student) read() {
	fmt.Println("%v is reading book", s.name)
}

func main() {
	var person Person
	person = Student{name: "zhangsan", age: 14}
	// 不能这样，接口只能访问自己的
	// person.name = "lisi"
	// 这样也不行，GoLang不支持这样子
	// student := Student(p)

	// 利用ok-idiom模型来还原接口到原来的子类
	// 可以用类型断言来获取，person.(Student)
	if student, ok := person.(Student); ok {
		student.name = "lisi"
		fmt.Println(student)
	}

	// 利用type switch还原
	switch s := person.(type) {
	case Student:
		s.name = "wangwu"
		s.age = 18
		fmt.Println(s)
	default:
		fmt.Println("错误的类型")
	}
}
