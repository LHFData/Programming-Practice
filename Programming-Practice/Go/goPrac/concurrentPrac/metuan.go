package main

import (
	"fmt"
	"time"
)
func work1(ch chan int){
	var i int
	i=<-ch
	fmt.Println(i)
}
func work2(ch chan int){
	i:=3
	ch<-i
	fmt.Println("input")
}
func main() {
	ch:=make(chan int)
	go work1(ch)
	go work2(ch)
	time.Sleep(10*time.Second)
	//a := 0
	//fmt.Scan(&a)
	//fmt.Printf("%d\n", a)
	fmt.Printf("Hello World!\n");
}
