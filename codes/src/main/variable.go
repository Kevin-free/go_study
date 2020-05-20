// 变量与内置数据类型

package main

import (
	"fmt"
	"os"
	"reflect"
)

// 首字母大写，是公有的
func Simple() {
	f, err := os.Open("infile")
	f, err = os.Create("outfile")
	fmt.Println(f, err)
}

func Point() {
	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // 1
	*p = 2          // equal to x = 2
	fmt.Println(x)  // 2
}

// 首字母小写，是私有的
func incr(p *int) int {
	*p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	return *p
}

func delta(old, new int) int {
	return new - old
}

// 测试字符串
func testStr() {
	str2 := "Go语言"
	runeArr := []rune(str2)
	// reflect.TypeOf().Kind() 可以知道某个变量的类型
	fmt.Println(reflect.TypeOf(runeArr[2]).Kind())
	fmt.Println(runeArr[2], string(runeArr[2]))
	fmt.Println("len(runeArr):", len(runeArr))
}

// 测试切片
func testSlice() {
	// 声明切片
	slice1 := make([]float32, 0)
	slice2 := make([]float32, 3, 5)
	fmt.Println(len(slice1), cap(slice2)) // 0 5

	// 使用切片
	// 添加元素，切片容量可以根据需要自动扩展
	slice2 = append(slice2, 1, 2, 3, 4)   // [0,0,0,1,2,3,4]
	fmt.Println(len(slice2), cap(slice2)) // 7 12
	// 自切片 [start, end]
	sub1 := slice2[3:] // [1 2 3 4]
	sub2 := slice2[:3] // [0 0 0]
	//sub3 := slice2[1:4] // [0 0 1]
	// 合并切片
	combined := append(sub1, sub2...) // sub2... 是切片解构的写法，将切片解构为N个独立的元素
	fmt.Println(combined)             // [1 2 3 4 0 0 0]
}

// 测试Map
func testMap() {
	// 仅声明
	m1 := make(map[string]int)
	// 声明时初始化
	m2 := map[string]string{
		"Tom":   "Male",
		"Jerry": "Female",
	}
	// 赋值/修改
	m1["Tom"] = 18
	fmt.Println(m2) // map[Jerry:Female Tom:Male]
}

// 测试指针
func testPoint() {
	str := "Golang"
	var p *string = &str
	*p = "hello"
	fmt.Println(str)
}

func fakeAdd(num int) {
	num += 1
}
func realAdd(num *int) {
	*num += 1
}
func testAdd() {
	num := 100
	fakeAdd(num)
	fmt.Println(num) //100

	realAdd(&num)
	fmt.Println(num) //101
}

func main() {
	//fmt.Println("===test incr===")
	//v := 1
	//incr(&v)
	//fmt.Println(incr(&v))
	//
	//res := Fib(3)
	//fmt.Println("fib res:",res)

	//testStr()

	//testSlice()

	//testMap()

	//testPoint()

	testAdd()
}
