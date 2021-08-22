package main

import (
	"fmt"
	"runtime"
)

func main(){
	n:=runtime.GOMAXPROCS(2)
	fmt.Printf("n=%d\n",n)
	for{
		go fmt.Print(0)
		fmt.Print(1)
	}
}
