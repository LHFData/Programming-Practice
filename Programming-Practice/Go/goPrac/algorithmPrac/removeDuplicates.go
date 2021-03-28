package main

import "fmt"
func removeDuplicates(nums []int) int {
	var p int
	var q int
	p=0
	q=1
	for q<len(nums) {
		if nums[q] != nums[p] {
			if(q-p<1) {
				p++
				q++
			}else{
				nums[p+1]=nums[q]
				p++
			}
		}else{
			q++
		}
	}
	return p+1
}
func main(){
	 input:=[]int{1,1,1,2,2,3,4,4,5,6,6}
	 fmt.Println(removeDuplicates(input))
}
