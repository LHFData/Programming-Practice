package main

import (
	"fmt"
	"time"
)

func consumer(ch chan int){
	num:=<-ch
	fmt.Println(num)
}
func producer(ch chan int){
	ch<-2
	time.Sleep(time.Second)
}
func main()  {
	ch:=make(chan int,2)
	go producer(ch)
	go consumer(ch)
	for{

	}
}
