// 结构体，方法

package main

import "fmt"

type Student struct {
	name string
	age  int
}

// (stu *Student) 说明这是 Student 结构体的方法
func (stu *Student) hello(person string) string {
	return fmt.Sprintf("hello %s, I am %s", person, stu.name) //hello Jerry, I am Tom
}

func main() {
	stu := &Student{
		name: "Tom",
	}
	msg := stu.hello("Jerry")
	fmt.Println(msg)

	// new 方式实例化
	stu2 := new(Student)
	fmt.Println(stu2.hello("Alice")) //hello Alice, I am  // name 被赋默认值""
}
