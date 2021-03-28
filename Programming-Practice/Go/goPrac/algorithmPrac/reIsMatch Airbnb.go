package main

import "fmt"

var pre uint8
/*func isMatch(s string,p string) bool{


	if len(s)==0&&p[0]!='*'{
		return false
	}
	scount:=0//待配串滑动距离
	pcount:=0
	for i:=0;i<len(p);i++{

		for j:=i+scount-pcount;j<len(s);j++{
			if p[i]==s[j]||p[i]=='.'{
				pre=s[j]
				break
			}else if p[i]=='*'&&s[j]==pre{

				if(j+1<len(s)&&s[j+1]!=pre){
					break
				}
				scount++
				continue
			}else if p[i]=='*'&&s[j]!=pre{
				pcount++//出现了适配后又丢失的*号，应该略过对应长度的待配式,否则影响for循环第一式
				break
			}else{
				return false
			}
		}

	}

	return true
}
*/
func isMatch(s string,p string)bool{
	m,n:=len(s),len(p)
	matches:=func(i,j int)bool{
		if i==0{
			return false
		}
		if p[j-1]=='.'{
			return true
		}
		return s[i-1]==p[j-1]
	}
	f:=make([][]bool,m+1)//二维数组，大小为被匹配串的长度
	for i:=0;i<len(f);i++{
		f[i]=make([]bool,n+1)
	}
	f[0][0]=true
	for i:=0;i<=m;i++{
		for j:=1;j<=n;j++{
			if p[j-1]=='*'{
				f[i][j]=f[i][j]||f[i][j-2]
				if matches(i,j-1){
					f[i][j]=f[i][j]||f[i-1][j]
				}
			}else if matches(i,j){
				f[i][j]=f[i][j]||f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}
func main(){
	//fmt.Println(isMatch("LHFF","..F*"))
	fmt.Println(isMatch("LHFFFFFCG","LHF*C*A*G"))
}