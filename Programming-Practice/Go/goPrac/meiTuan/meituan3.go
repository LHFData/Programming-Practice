package main

import (
	"fmt"
)
func getRes3(s,a string)int{
	var res int
	m:=make(map[byte]int)
	for i:=0;i<len(s);i++{
		_,ok:=m[s[i]]
		if !ok{
			m[s[i]]=i
		}
	}
	pre:=-1
	for i:=0;i<len(a);i++{
		tmp,ok:=m[a[i]]
		if ok{
			if pre==-1{
				res+=tmp
				pre=tmp
				continue
			}
			if pre>=tmp{
				res=res+(len(s)-1-pre)+tmp
			}else{
				res=res+tmp-pre-1
			}
			pre=tmp
		}else{
			return -1
		}
	}
	return res
}
func main(){
	var stream string
	fmt.Scan(&stream)
	var aim string
	fmt.Scan(&aim)
	fmt.Println(getRes3(stream,aim))
}
