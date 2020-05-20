// 并发编程

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var ch = make(chan string, 10) // 创建大小为10的缓冲信道

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second) // 模拟耗时操作
	wg.Done()               // 为 wg 减去一个计数
}

func downloadByChannel(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second) // 模拟耗时操作
	ch <- url               // 将 url 发给信道
}

func main() {
	//for i := 0; i < 3; i++ {
	//	wg.Add(1) // 为 wg 添加一个计数
	//	go download("a.com/"+string(i+'0')) // 启动新的协程（goroutine）并发执行download函数
	//}
	//wg.Wait() // 等待所有的协程执行结束
	//fmt.Println("Done!")

	for i := 0; i < 3; i++ {
		go downloadByChannel("a.com/" + string(i+'0')) // 启动新的协程（goroutine）并发执行download函数
	}
	for i := 0; i < 3; i++ {
		msg := <-ch // 等待信道返回消息
		fmt.Println("finish", msg)
	}
	fmt.Println("Channel Done!")
}
