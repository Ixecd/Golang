package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func testChannle() {
// 	ch1 := make(chan<- int, 10) // 只写信道
// 	ch2 := make(<-chan int, 10) // 只读信道
// 	ch3 := make(chan int, 10) // 双向信道
// 	// 单向通道最主要地用途就是约束其他代码的行为
// }

func SendInt(ch chan<- int) {
	ch <- rand.Intn(1000)
}

func RecvInt(ch <-chan int) {
	gt := <-ch
	fmt.Println("get value:", gt)
}

func main() {

	ch := make(chan int, 10)
	go SendInt(ch)
	go RecvInt(ch) // 如果go RecvInt(ch)终端并不会显示, 原因是Main结束了
	time.Sleep(1 * time.Second)
}