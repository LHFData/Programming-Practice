package main

import "fmt"

func canJump(nums []int) bool {

	res:=false

	var traverse func(idx int)
	traverse=func(idx int){
		if idx==0{
			res=true
			fmt.Println(res)
			return
		}
		traverse(idx-1)
	}
	traverse(3)
	return res
}
func main(){
	input:=[]int{2,1,3,3,4}
	fmt.Println(canJump(input))
}
