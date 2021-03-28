package main

import "fmt"

func minimumTotal(triangle [][]int) int{
	level:=len(triangle)
	var dp [][]int
	dp=make([][]int,level)
	for i:=level-1;i>=0;i--{
		var line []int
		line=make([]int,i+1)
		for j:=0;j<i+1;j++{
			if i==level-1{
				line[j]=triangle[i][j]
				continue
			}
			if dp[i+1][j]<dp[i+1][j+1]{
				line[j]=dp[i+1][j]+triangle[i][j]
			} else{
				line[j]=dp[i+1][j+1]+triangle[i][j]
			}
		}
		dp[i]=line
	}
	return dp[0][0]

}
func main(){
	a:=[][]int{{-1},{2,3},{1,-1,-3}}
	fmt.Println(minimumTotal(a))
}