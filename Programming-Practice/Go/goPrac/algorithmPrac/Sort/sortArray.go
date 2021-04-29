package main

import "fmt"

func quickSort(nums []int,start int ,end int){
	if start<end {
		left := start
		right := end

		temp := nums[start]
		//找到哨兵
		for left < right {

			for nums[right] > temp && left < right {
				right--
			}
			if left < right {
				nums[left] = nums[right]
				left++
			}
			//right覆盖left
			for nums[left] < temp && left < right {
				left++
			}
			//left覆盖right
			if left < right {
				nums[right] = nums[left]
				right--
			}

		}
		nums[left] = temp
		quickSort(nums, start, left-1)
		quickSort(nums, left+1, end)
	}
}
func sortArray(nums []int) []int {
	temp:=nums
	quickSort(temp,0,len(nums)-1)
	return temp
}

func main(){
	n:=[]int{-2,3,-5}
	sortArray(n)
	fmt.Println(n)
}