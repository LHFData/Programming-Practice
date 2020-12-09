package main

import (
	"fmt"
	"time"
)

func main(){
	for t:=range time.Tick(time.Second){
		fmt.Println("tick at ",t)
	}
}