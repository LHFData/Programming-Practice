package main

import (
	"fmt"
	"time"
)

func printString(value string) {
	for i := 0; i < 10; i++ {
		fmt.Println(value)
		time.Sleep(time.Second)
	}
}
func main() {
	go printString("A")
	go printString("C")

	time.Sleep(time.Second * 10)
}
