package main

import (
	"fmt"
	"sync"
	"time"
)
func thrWorker(wg *sync.WaitGroup){

	fmt.Println("3号完成")
	wg.Done()
}
func main(){
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		time.Sleep(2*time.Second)
		fmt.Println("1号完成")
		wg.Done()
	}()
	go func() {
		time.Sleep(5*time.Second)
		fmt.Println("2号完成")
		wg.Done()
	}()
	go thrWorker(&wg)
	wg.Wait()
	fmt.Println("收工")
}
