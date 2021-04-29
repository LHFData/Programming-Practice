package main

func strStr(haystack string, needle string) int {
	if len(haystack)<len(needle){
		return -1
	}
	if len(needle)==0{
		return 0
	}
	index:=-1
	for i:=0;i<len(haystack);i++{
		if haystack[i]==needle[0]{
			if len(needle)==1{
				index=i
				break
			}
			flag:=true
			for j:=1;j<len(needle)&&i+j<len(haystack);j++{
				if haystack[i+j]!=needle[j]{
					flag=false
					break
				}
			}
			if i+len(needle)>len(haystack){
				index=-1
				break
			}
			if flag{
				index=i
				break
			}

		}
	}
	return index
}
