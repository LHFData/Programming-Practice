package main

import "fmt"


func removeElement(nums []int, val int) int {
	order:=0
	unorder:=0
	for unorder<len(nums){
		if nums[unorder]==val{
			unorder++
			continue
		}
		nums[order]=nums[unorder]
		order++
		unorder++
	}
	return order
}

func main(){
	a:=[]int{1,3,2,3,1,3}
	fmt.Println(removeElement(a,3))
}
