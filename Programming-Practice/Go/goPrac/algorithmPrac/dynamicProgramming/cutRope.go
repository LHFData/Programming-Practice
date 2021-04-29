package main
//动规能做但是要记住一个结论，尽可能分为3作为一份是最大的，如果后面有不能分为3的部分，如2，1，0，2则保留2，1则与前一个3结合成为2，2，0则最优
func cuttingRope(n int) int {


	dp:=make([]int,n+1)
	dp[0]=0
	dp[1]=1
	dp[2]=1
	max:=func(a int,b int)int{
		if a>b{
			return a
		}
		return b
	}
	for i:=3;i<=n;i++{
		for j:=2;j<=i;j++{
			dp[i]=max(dp[i],max(j*(i-j),j*dp[i-j]))
		}
	}
	return dp[n]
}
