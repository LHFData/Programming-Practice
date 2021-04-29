package main

import (
	"fmt"
)

func fuck(n [][]int){
	n[0][0],n[0][1]=n[0][1],n[0][0]
}
func change(n []int){
	n[0],n[1]=n[1],n[0]
}
func main()  {
	s:=[]int{1,2,3,4,5}
	ss:=[][]int{{1,2},{3,4}}

	fuck(ss)
	//change(s[:2])
	//change(s[2:4])
	fmt.Println(s[1:])

}
