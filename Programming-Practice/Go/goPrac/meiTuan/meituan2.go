package main

import (
	"fmt"
)

//全排列，大数
func getRes2(a []int,b []int) int{
	var res int
	var curBookIndex int
	length:=len(a)
	judge:=make([]bool,length)
	var recurse func(idx int)
	recurse=func(idx int){
		if curBookIndex==length{
			res++
			fmt.Println(judge)
			return
		}
		if judge[idx]{
			return
		}
		if a[curBookIndex]<=b[idx]{
			curBookIndex++
			judge[idx]=true
			for i:=0;i<length;i++{

				recurse(i)
			}
			judge[idx]=false
			curBookIndex--
		}

	}
	recurse(0)

	return res
}
func main(){
	var length int
	fmt.Scan(&length)
	var a []int
	var b []int

	for i:=0;i<length;i++{
		tmp:=0
		fmt.Scan(&tmp)
		a=append(a,tmp)
	}
	for i:=0;i<length;i++{
		tmp:=0
		fmt.Scan(&tmp)
		b=append(b,tmp)
	}
	fmt.Println(getRes2(a,b))
}
