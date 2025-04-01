package main

import "fmt"

func main() {
	// testStruct()
	// testStructPoint()
	testStructMethod()
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

func testStructPoint() {
	type Person struct {
		name string
		age  int
	}

	p := &Person{}
	// 这里由于"."的优先级比"*"高，所以需要加上括号；不过这种写法太繁琐了，GoLang提供了语法糖，直接用p.age即可
	(*p).name = "jackpot"
	p.age = 24
	fmt.Println(p)
}

type BoyNextDoor struct {
	name string
	age  int
}

// 这种写法不会更新传递的对象，值传递，在调用这个方法的时候，会拷贝一个对象的副本
func (b BoyNextDoor) setName(name string) {
	b.name = name
}

// 这样写才会更新传递的对象，引用传递
func (b *BoyNextDoor) setAge(age int) {
	b.age = age
}

// 自己定义的类型就能实现方法
type integer int

func (i integer) add() {
	fmt.Println("integer add invoke")
}

func testStructMethod() {
	boyNextDoor := BoyNextDoor{name: "zhangsan", age: 18}
	fmt.Println("boyNextDoor = ", boyNextDoor)
	boyNextDoor.setName("lisi")
	fmt.Println("boyNextDoor = ", boyNextDoor)
	// 所有的setAge都可以生效，这几种写法实测都是可以的
	boyNextDoor.setAge(20)
	fmt.Println("boyNextDoor = ", boyNextDoor)
	boyNextDoorPtr := &boyNextDoor
	boyNextDoorPtr.setAge(22)
	fmt.Println("boyNextDoor = ", boyNextDoor)
	(*boyNextDoorPtr).setAge(24)
	fmt.Println("boyNextDoor = ", boyNextDoor)

	var i integer = 6
	i.add()
}
