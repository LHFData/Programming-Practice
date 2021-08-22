package main

import (
	"fmt"
)
type lr struct {
	left *lr
	val int
	right *lr
}
func fuck(n [][]int){
	n[0][0],n[0][1]=n[0][1],n[0][0]
}
func change(n []int){
	n[0],n[1]=n[1],n[0]
}
func main()  {
	s:=[]int{1,2,3,4,5}
	ss:=[][]int{{1,2},{3,4}}
	l:=new(lr)
	var ll []*lr
	l.val=3
	ll=append(ll,l.right)
	ll=append(ll,l.left)
	h:=lr{val: 3,left: nil,right: nil}
	ll[0]=&h
	fuck(ss)
	//change(s[:2])
	//change(s[2:4])
	fmt.Println(len(s[:5]))
	fmt.Println(2<<1)
}
