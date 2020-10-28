package main

import "fmt"

func main3(){
	list:=*new([]int)//取指针内容的引用就无所谓了
	//list:=new([]int) //这样写就会过不了编译
	list=append(list,1)//一个指向数组的指针被传入了，自然过不了
	fmt.Println(list)
}
var(
	size=1024//不能用:= var size=1024已足够完整。
	max_size=size*2
)
func main(){

	fmt.Println(size,max_size)
	s1:=[]int{1,2,3}
	s2:=[]int{4,5}
	s1=append(s1,s2...)//这个是切片合并的简单写法，...是三个英文句号
	fmt.Println(s1)

}