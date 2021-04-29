package main

func longestCommonPrefix(strs []string) string {
	var res string
	_=res
	if len(strs)==0{
		return ""
	}
	if len(strs)==1{
		return strs[0]
	}
	shortestStr:=len(strs[0])
	shortestIndex:=0
	for i:=0;i<len(strs);i++{

		if len(strs[i])<shortestStr{
			shortestStr=len(strs[i])
			shortestIndex=i
		}
		if len(strs[i])==0{
			return res
		}
	}
	for i:=0;i<shortestStr;i++{
		current:=strs[0][i]
		flag:=true
		for j:=1;j<len(strs);j++{
			if strs[j][i]==current{
				continue
			}else{
				flag=false
				break
			}
		}
		if !flag{
			return strs[shortestIndex][0:i]

		}
	}
	res=strs[shortestIndex]
	return res
}



