package main

import "fmt"
//暴力解法
//func searchInsert(nums []int, target int) int {
//
//	for i:=0;i<len(nums)-1;i++{
//		if nums[i]==target{
//			return i
//		}else if nums[i]<target&&target<nums[i+1]{
//			return i+1
//		}
//
//	}
//	if target==nums[len(nums)-1]{
//		return len(nums)-1
//	}else if target>nums[len(nums)-1]{
//		return len(nums)
//	}else if target<nums[0]{
//		return 0
//	}
//
//	return -1
//}
func searchInsert(nums []int, target int)int{
	left:=0
	right:=len(nums)
	middle:=(left+right)/2
	if (target>nums[right-1]){
		return right
	}
	for left<right{
		middle=(left+right)/2
		if(nums[middle]>target){
			right=middle
			continue
		}else if(nums[middle]<target){
			left=middle+1
			continue
		}
	}
	return middle-1
}
func main()  {
	a:=[]int{1,2,6,7}
	fmt.Println(searchInsert(a,3))
}
