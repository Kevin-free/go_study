// 流程控制

package main

import "fmt"

func testIfElse() {
	age := 18
	if age < 18 {
		fmt.Println("kid")
	} else {
		fmt.Println("adult")
	}

	// 可以简写为
	//if age:= 18; age < 18 {
	//	fmt.Println("kid")
	//}else {
	//	fmt.Println("adult")
	//}
}

func testSwitch() {
	type Gender int8
	const (
		MALE   Gender = 1
		FEMALE Gender = 2
	)

	gender := MALE

	switch gender {
	case MALE:
		fmt.Println("male")
	case FEMALE:
		fmt.Println("female")
	default:
		fmt.Println("unknown")
	}
	// 无需 break
	// male
}

func testFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		if sum > 3 {
			break
		}
		sum += i
	}
	fmt.Println(sum)
}

// 对数组（arr），切片（slice），字典（map）使用for range 遍历
func testFor2() {
	// 数组的for循环遍历
	nums := []int{1, 2, 3, 4}
	for i, num := range nums {
		fmt.Println(i, num)
	}
	//0 1
	//1 2
	//2 3
	//3 4

	// 字典的for循环遍历
	m2 := map[string]string{
		"Tom":   "Male",
		"Jerry": "Female",
	}
	for key, value := range m2 {
		fmt.Println(key, value)
	}
	//Tom Male
	//Jerry Female
}

func main() {
	//testIfElse()
	//testSwitch()
	//testFor()
	testFor2()
}
