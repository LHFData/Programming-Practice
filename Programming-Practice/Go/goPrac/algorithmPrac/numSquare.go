package main

func numSquare(n int) int{
	dp:=make([]int,n+1)
	Min:=func(a int ,b int ) int{
		if a<=b{
			return a
		}else{
			return b
		}
	}
	for i:=1;i<=n;i++{
		dp[i]=i
		for j:=1;i-j*j>0;j++{
			dp[i]=Min(dp[i],dp[i-j*j]+1)
		}
	}
	return dp[n]

}
func main(){
	numSquare(13)
}
