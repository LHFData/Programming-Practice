package main

import "fmt"

func isMatch(s string, p string) bool {
	dp:=make([][]bool,len(s)+1)//比长度大一
	match:=func (i int,j int) bool{//匹配的不是当前，是前面
		if i==0{//空出一行
			return false
		}
		if p[j-1]=='.'{//前一个字符为.则匹配所有s[i]
			return true
		}
		return s[i-1]==p[j-1] //取决于前串是否匹配
	}
	for i:=0;i<len(s)+1;i++{
		dp[i]=make([]bool,len(p)+1)

	}
	dp[0][0]=true //空串也匹配
	for i:=0;i<len(s);i++{
		for j:=1;j<len(p);j++{
			if p[j-1]=='*'{//p的前一个是星号
				dp[i][j]=dp[i][j]||dp[i][j-2]//那么当前的值取决于当前是否匹配或者当前的s[i]与当前的j是否匹配，否则看是否与*之前的的串匹配
				if match(i,j-1){//看i-1和j-2是否匹配，即当前的前一个是否和*之前匹配
					dp[i][j]=dp[i-1][j]||dp[i-2][j]//
				}
			}else if match(i,j){
				dp[i][j]=dp[i][j]||dp[i-1][j-1]
			}
	}
	}
	return dp[len(s)-1][len(p)-1]
}
func main (){
	fmt.Println(isMatch("AA","A*"))
}

