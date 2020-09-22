package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var len = len(nums)
	var m = make(map[int]int)
	for i := 0; i < len; i++ {
		m[nums[i]] = i
		fmt.Printf("%d,%d \n", nums[i], i)
	}
	for k, v := range nums {
		var sear = target - v
		var num1 = k
		var num2 = 0
		num2, ok := m[sear]
		fmt.Printf("%d,%d \n", sear, k)
		if ok {
			if num1 != num2 {
				var r = []int{num1, num2}
				return r
			}
		}
	}
	var r = []int{0, 0}
	return r
}

func main() {
	//两数之和
	//var nums = []int{3,2,4}
	//fmt.Print(twoSum(nums, 6))

	//翻转整数
	var nums int = -655351
	fmt.Print(reverse(nums))
}
