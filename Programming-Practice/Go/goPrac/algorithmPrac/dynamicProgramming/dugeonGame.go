package main

import "fmt"

func calculateMinimumHP(dungeon [][]int) int {

	var dp [][]int

	min:=func(a int,b int)int{
		if a<b{
			return a
		}
		return b
	}

	m:=len(dungeon)-1
	n:=len(dungeon[0])-1
	for i:=0;i<=m;i++{
		var temp []int
		for j:=0;j<=n;j++{
			temp=append(temp,0)
		}
		dp=append(dp,temp)
	}
	//初始化DP
	for i:=m;i>=0;i--{
		for j:=n;j>=0;j--{
			if i+1>m&&j+1>n{
				if dungeon[i][j]>=0{
					dp[i][j]=1
				}else{
					dp[i][j]=1-dungeon[i][j]
				}
				continue
			}
			if i+1>m{
				if dungeon[i][j]>=0{
					temp:=dungeon[i][j]-dp[i][j+1]
					if temp>=0{
						dp[i][j]=1
					}else {
						dp[i][j]=-temp
					}
				}else{
					if dp[i][j+1]==0{
						dp[i][j]=1-dungeon[i][j]
						continue
					}
					dp[i][j]=-dungeon[i][j]+dp[i][j+1]
				}
				continue
			}
			if j+1>n{
				if dungeon[i][j]>=0{
					temp:=dungeon[i][j]-dp[i+1][j]
					if temp>=0{
						dp[i][j]=1
					}else {
						dp[i][j]=-temp
					}
				}else{
					if dp[i+1][j]==0{
						dp[i][j]=1-dungeon[i][j]
						continue
					}
					dp[i][j]=-dungeon[i][j]+dp[i+1][j]
				}
				continue
			}
			if dungeon[i][j]>=0 {
				temp:=dungeon[i][j]-min(dp[i][j+1], dp[i+1][j])
				if temp>=0{
					dp[i][j]=1
				}else{
					dp[i][j]=-temp
				}
			}else{
				if min(dp[i][j+1],dp[i+1][j+1])==0{
					dp[i][j]=-dungeon[i][j]+1
					continue
				}
				dp[i][j]=-dungeon[i][j]+min(dp[i][j+1],dp[i+1][j])
			}


		}
	}
	if dp[0][0]==0{
		return 1
	}
	return dp[0][0]
}
func main(){
	fmt.Println(calculateMinimumHP([][]int{{3,0,-3},{-3,-2,-2},{3,1,-3}}))
}