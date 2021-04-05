package main

func uniquePaths(m int, n int) int {
	//初始化
	dp:=make([][]int,m)
	for i:=0;i<m;i++{
		temp:=make([]int,n)
		dp[i]=temp
	}
	dp[0][0]=1
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if i==0&&j==0{
				continue
			}
			if i-1<0{
				dp[i][j]=0+dp[i][j-1]
			}else if j-1<0{
				dp[i][j]=dp[i-1][j]
			}else{
				dp[i][j]=dp[i-1][j]+dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]

}
