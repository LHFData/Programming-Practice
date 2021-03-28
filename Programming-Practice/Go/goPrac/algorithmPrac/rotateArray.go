package main

import "fmt"

func rotate(nums []int, k int)  {
	//暴力，链式存放，多次翻转，三种方案
	//链式存放会在数组长为k的倍数时原地打转
	//翻转比较靠谱
	swap:=func(n []int,start int,end int){
		for start<end {
			temp := n[start]
			n[start] = n[end]
			n[end] = temp
			start++
			end--
		}
	}
		k=k%len(nums)
		swap(nums,0,len(nums)-1)
		swap(nums,0,k-1)
		swap(nums,k,len(nums)-1)
	fmt.Println(nums)
}
func main()  {
	n:=[]int{1}
	rotate(n,2)

}
