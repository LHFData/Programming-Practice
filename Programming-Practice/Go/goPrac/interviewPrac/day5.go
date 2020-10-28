package main

import "fmt"

func main(){
	sn1:=struct{
		age int
		name string
	}{age:11,name:"qq"}
	sn2:=struct{
		age int
		name string
	}{age:11,name:"qq"}
	if sn1==sn2{
		fmt.Println("sn1==sn2")
	}
	sm1:=struct{
		age int
		m map[string]string
	}{age:11,m:map[string]string{"a":"1"}}
	sm2:= struct {
		age int
		m map[string]string
	}{age:11,m:map[string]string{"a":"1"}}
	if sm1==sm2{
		fmt.Println("sm1==sm2")// bool,数值，字符，指针，数组都是可比较的，而切片，函数，map都是不可比较的
	}
}
