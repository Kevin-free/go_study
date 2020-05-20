// 函数

package main

import (
	"fmt"
	"os"
)

func add(num1 int, num2 int) int {
	return num1 + num2
}

func div(num1 int, num2 int) (int, int) {
	return num1 / num2, num1 % num2
}

func testError() {
	// 下划线_的意思是忽略这个变量
	// go中如果声明变量而不使用的话，编译器会报错
	_, err := os.Open("filename.txt") // Open 返回值两个（*File, error）
	if err != nil {
		fmt.Println(err)
	}
}

func testUnknownError(index int) int {
	arr := [3]int{1, 2, 3}
	return arr[index]
}

// 捕捉异常，相当于Java中的 try-catch
func deferError(index int) (ret int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Some error happened!", r)
			ret = -1
		}
	}()
	arr := [3]int{2, 3, 4}
	return arr[index]
}

func main() {
	quo, rem := div(10, 3)
	fmt.Println(quo, rem)   // 3 1
	fmt.Println(add(10, 7)) // 17

	//fmt.Println(testUnknownError(5))
	//panic: runtime error: index out of range
	//goroutine 1 [running]:
	//Process finished with exit code 2

	fmt.Println(deferError(5))
	//Some error happened! runtime error: index out of range
	//-1
	//finished
	fmt.Println("finished")

}
