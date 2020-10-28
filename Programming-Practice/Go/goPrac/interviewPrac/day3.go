package main

import "fmt"
//切片带长度初始化后，初值全是0且已分配
func main1(){
	s:=make([]int,5)
	s = append(s, 1,2,3)
	fmt.Println(s)
}
func main2(){
	s:=make([]int,0)
	s=append(s,1,2,3,4)
	fmt.Println(s)
}


/* 返回值必定义类型
func Mui(x,y int )(sum int,error){
	return x+y,nil
}*/
type test struct{
	val int
	prev *test
}
func main(){
	main1()
	main2()
	testNew:=new(test)//new通常用来定义值类型，返回的是个地址，*test
	testMake:=make([]int,1)//返回的是一个引用
	testNew.val=2
	testNew.prev=nil
	fmt.Println("new:",testNew)
	fmt.Println("make:",testMake)

}