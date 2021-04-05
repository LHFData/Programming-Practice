package main

import "fmt"
func longestPalindrome(s string) string {
	//isPalindrome:= func(sub string) bool{
	//	length:=len(sub)
	//	if length%2==1{
	//		//奇数长度
	//		for i:=0;i<length/2;i++{
	//			if sub[i]==sub[length-1-i]{
	//				continue
	//			}else{
	//				return false
	//			}
	//		}
	//	}else{
	//		//偶数长度
	//		for i:=0;i<length/2;i++{
	//			if sub[i]==sub[length-1-i]{
	//				continue
	//			}else{
	//				return false
	//			}
	//		}
	//	}
	//	return true
	//}
	var res string
	dp:=make([][]bool,len(s))
	for i:=0;i<len(s);i++{
		dp[i] =make([]bool,len(s))
	}
	dp[0][0]=true
	longestlen:=0

	for l:=0;l<len(s);l++{
		for i:=0;i+l<len(s);i++{
			if l>2{
				if s[i]==s[i+l] {
					dp[i][i+l] = dp[i+1][i+l-1]
				}else{
					dp[i][i+l]=false
				}
			}else{
				if s[i]==s[i+l]{
					dp[i][i+l]=true

				}else{
					dp[i][i+l]=false
				}
			}
			if dp[i][i+l]&&l>=longestlen{
				res=s[i:i+l+1]
				longestlen=l
			}
		}
	}
	for i:=0;i<len(dp);i++{
		fmt.Println(dp[i])
	}
	fmt.Println(res)

	return res
}

func main()  {
	longestPalindrome("a")
}