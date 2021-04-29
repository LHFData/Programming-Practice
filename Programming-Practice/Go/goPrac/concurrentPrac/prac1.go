package main

import (
	"fmt"
	"time"
)
func running(id int,limit int){
	var times int
	for {
		times++
		fmt.Printf("id:%d tick:%d\n",id,times)
		time.Sleep(time.Second)
		if times==limit{
			break
		}
	}
}
func main(){

	var input string
	go running(1,5)
	go running(2,6)
	fmt.Scanln(&input)
}
