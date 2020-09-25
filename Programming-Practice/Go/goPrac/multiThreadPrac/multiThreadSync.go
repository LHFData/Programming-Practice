package main

import (
	"fmt"
	"sync"
	"time"
)

// 线程通信，共享数据方式
//共享变量 counter
var counter int = 0

func Count(lock *sync.Mutex, id int) {
	lock.Lock()
	counter++
	fmt.Println(counter, "I'm thread", id)
	lock.Unlock()
}
func main() {
	lock := &sync.Mutex{}
	for i := 0; i < 5; i++ {
		go Count(lock, i)
	}
	time.Sleep(time.Second * 5) //主线程挂起，让子线程跑完
}
