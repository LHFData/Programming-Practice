package main

import (
	"fmt"
	"time"
)

func WriteIn(ch chan int) {
	for i := 11; i < 20; i++ {
		ch <- i //通道输入
	}
}

func ReadOut(ch chan int) {
	for i := 0; i < 10; i++ {
		value := <-ch //通道输出
		fmt.Println(value)
	}
}
func main() {
	ch := make(chan int)
	go WriteIn(ch)
	go ReadOut(ch)
	time.Sleep(time.Second * 5)
}
