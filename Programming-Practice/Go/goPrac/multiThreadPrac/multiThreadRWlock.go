package main

import (
	"fmt"
	"sync"
	"time"
)

type MutexRWInfo struct {
	mutex sync.RWMutex
	infos []int
}

func (m *MutexRWInfo) writeInfo(value int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	fmt.Println("write start", value)
	m.infos = append(m.infos, value)
	fmt.Println("write start end", value)
}
func (m *MutexRWInfo) readInfo(value int) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	fmt.Println("read start", value)
	fmt.Println("read end", value)
}

//下面试一下条件变量
func printIntValue(value int, cond *sync.Cond) {
	cond.L.Lock()
	if value < 5 {
		cond.Wait()
	}
	fmt.Println("cond output", value)
	cond.L.Unlock()
}
func main() {
	m := MutexRWInfo{}
	for i := 0; i < 10; i++ {
		go m.writeInfo(i)
		go m.readInfo(i)
	}
	time.Sleep(time.Second * 3)
	fmt.Println(m.infos)

	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)
	for i := 0; i < 10; i++ {
		go printIntValue(i, cond)
	}
	time.Sleep(time.Second * 1)
	cond.Signal()
	time.Sleep(time.Second * 1)
	cond.Broadcast()
	time.Sleep(time.Second * 1)
}
