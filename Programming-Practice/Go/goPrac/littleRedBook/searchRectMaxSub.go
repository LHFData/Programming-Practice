
package main
import (
"fmt"
)

func main(){
	var n int
	fmt.Scan(&n)
	var score [][]int
	for i:=0;i<n;i++{
		var subScore []int
		for j:=0;j<n;j++{
			var tmp int
			fmt.Scan(&tmp)
			subScore=append(subScore,tmp)
		}
		score=append(score,subScore)
	}

	var dp [][]int
	for i:=0;i<n+1;i++{
		tmp:=make([]int,n+1)
		dp=append(dp,tmp)
	}
	var max int
	max=score[0][0]
	for i:=1;i<n+1;i++{
		for j:=1;j<n+1;j++{
			dp[i][j]=score[i-1][j-1]+dp[i][j-1]+dp[i-1][j]-dp[i-1][j-1]
			if dp[i][j]>max{
				max=dp[i][j]
			}
		}
	}
	fmt.Println(max)




}
