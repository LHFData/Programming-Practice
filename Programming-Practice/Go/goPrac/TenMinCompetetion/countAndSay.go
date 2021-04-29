package main

import (
	"fmt"
	"strconv"
)

func description(n string) string{
	var res string
	var mark string
	count:=1
	if len(n)==1{
		return strconv.Itoa(1)+n
	}
	for i:=0;i<len(n);i++{
		mark=string(n[i])
		if i==len(n)-1{
			if n[i]==n[i-1]{

				res=res+strconv.Itoa(count)+mark
				break
			}else{
				res=res+strconv.Itoa(1)+string(n[i])
				break
			}
		}
		if n[i]==n[i+1]{
			count++
		}else{
			res=res+strconv.Itoa(count)+mark
			count=1
		}
	}

	return res
}
func countAndSay(n int) string {
	var res string
	res="1"
	for i:=1;i<n;i++{
		res=description(res)
	}
	return res
}
func main(){
	fmt.Println(description("1"))
}
