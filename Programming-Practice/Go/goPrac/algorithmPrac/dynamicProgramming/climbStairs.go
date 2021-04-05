package main

import "fmt"
func climbStairs(n int) int {
	dp:=make([]int,n+1)

	if n<2{
		return 1
	}
	dp[0]=1
	dp[1]=1
	for i:=2;i<n+1;i++{
		dp[i]=dp[i-1]+dp[i-2]
	}
	return dp[n]
}
func main()  {
	fmt.Println(climbStairs(5))
}
