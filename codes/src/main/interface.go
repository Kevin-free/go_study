// 接口

package main

import "fmt"

type Person interface {
	getName() string
}

type Student2 struct {
	name string
	age  int
}

func (stu *Student2) getName() string {
	return stu.name
}

type Worker struct {
	name   string
	gender string
}

func (w *Worker) getName() string {
	return w.name
}

func main() {
	var p Person = &Student2{
		name: "Tom",
		age:  18,
	}
	fmt.Println(p.getName()) // Tom

	var _ Person = (*Worker)(nil) // 检测某个类型是否实现里某个接口的所有方法
}
