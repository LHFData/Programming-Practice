package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m:=make(map[int]int)
	for i:=0;i<len(nums);i++{
		_,ok:=m[nums[i]]
		if ok{
			return false
		}else{
			m[nums[i]] = nums[i]
		}

	}
	return true
}
func main()  {
	a:=[]int{1,3,4}
	fmt.Println(containsDuplicate(a))

}