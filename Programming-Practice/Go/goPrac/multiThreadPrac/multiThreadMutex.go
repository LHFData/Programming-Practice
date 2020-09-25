package main

import (
	"fmt"
	"sync"
	"time"
)

type MutexInfo struct {
	mutex sync.Mutex
	infos []int
}

func (m *MutexInfo) addInfo(value int) {
	m.mutex.Lock()
	m.infos = append(m.infos, value)
	m.mutex.Unlock()
}
func main() {
	m := MutexInfo{}
	for i := 0; i < 10; i++ {
		go m.addInfo(i + 10)
		go m.addInfo(i + 20)
	}
	time.Sleep(time.Second * 5)
	fmt.Println(m.infos)

}
