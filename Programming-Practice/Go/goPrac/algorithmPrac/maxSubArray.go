package main

import "fmt"

func maxSubArray(nums []int) int {
	//动态规划和分治法，选一个做，动态规划的考虑方式
	//若前面一个最大和对当前产生负增益，则弃置，否则在前面最大和和当前值之间取最大值
	Max:=func(m []int) int{
		max:=m[0]
		for i:=0;i<len(m);i++{
			if max<m[i]{
				max=m[i]
			}
		}
		return max
	}
	var dp []int
	dp=append(dp,nums[0])
	for i:=1;i<len(nums);i++{
		if nums[i]+dp[i-1]>nums[i]{
			dp=append(dp,dp[i-1]+nums[i])
		}else{
			dp=append(dp,nums[i])
		}
	}
	return Max(dp)
}

func main()  {
	a:=[]int{1,-2,1,3,4,-1,-2,3,6,9,1}
	fmt.Println(maxSubArray(a))
}