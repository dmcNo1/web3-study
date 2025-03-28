package main

import "fmt"

func main() {
	testStruct()
}

func testStruct() {
	type Person struct {
		name string
		age  int
	}
	type Class struct {
		name        string
		headTeacher string
	}
	// 结构体中的变量都是匿名属性，基础类型不能重复
	type Student struct {
		*Person
		Class
		int
	}

	person := Person{name: "zhangsan", age: 13}
	class := Class{name: "初一2班", headTeacher: "CD Xie"}
	student := Student{Person: &person, Class: class}
	// 匿名属性可以直接通过类型作为名称访问，如果同名，则需要指定对应的结构体类型
	fmt.Printf("student.name = %v, class.name = %v, student.int = %v\n", student.Person.name, student.Class.name, student.int)
}
