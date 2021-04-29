package main

import "fmt"

func permutation(nums []int,re []int,res *[][]int) {
	if len(nums)==1{
		re=append(re,nums[0])
		*res=append(*res,re)
		return
	}else{
		for i:=0;i<len(nums);i++{
			var r []int
			r=append(re,nums[i])
			var next []int
			next=make([]int,len(nums))
			copy(next,nums)
			next=append(next[:i],next[i+1:]...)
			permutation(next,r,res)
		}
	}
}
func permute(nums []int) [][]int {
	var res [][]int
	var re []int
	permutation(nums,re,&res)

	return res
}
func main(){
	fmt.Println(permute([]int{1,1,3}))
}
