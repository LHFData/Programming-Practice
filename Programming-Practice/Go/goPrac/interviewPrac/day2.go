package main

import "fmt"

func main(){
	slice:=[]int{0,1,2,3}
	m:=make(map[int]*int)
//for range 创建的是每个元素的副本而非引用，因此取来取去都会只拿到作为临时容器的val的地址。
	for key,val:=range slice{
		value:=val //创建新变量
		m[key]=&value
	}
	for k,v :=range m{
		fmt.Println(k,"->",v,"->",*v)
	}
}
