package main

import (
	"fmt"
	"runtime"
)

func main(){
	go func(s string) {
		for i:=0;i<2;i++{
			fmt.Println(s)
		}
	}("world")
	for i:=0;i<2;i++{
		fmt.Println("hello")
		runtime.Gosched()

	}
}
