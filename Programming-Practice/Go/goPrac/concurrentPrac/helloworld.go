package main

import (
	"fmt"
	"time"
)

func thread1(){
	fmt.Println("Hello world")

}
func main()  {
	go thread1()
	time.Sleep(1 * time.Second)

}