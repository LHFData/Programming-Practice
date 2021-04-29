package main

import "fmt"

func permu(nums []int,res *[][]int,re []int){
	if len(nums)==1{
		re=append(re,nums[0])
		*res=append(*res,re)
		return
	}
	count:=make(map[int]bool)
	for i:=0;i<len(nums);i++{
		//剪枝
		_,ok:=count[nums[i]]
		if ok{

			continue
		}

		count[nums[i]]=true
		var r []int
		var next []int
		r= append(re, nums[i])
		next=make([]int,len(nums))
		copy(next,nums)
		next=append(next[:i],next[i+1:]...)
		permu(next,res,r)
	}

}

func permuteUnique(nums []int) [][]int{
	var res [][]int
	var r [] int
	permu(nums,&res,r)
	return res
}
func main(){
	fmt.Println(permuteUnique([]int{1,1,3}))
}