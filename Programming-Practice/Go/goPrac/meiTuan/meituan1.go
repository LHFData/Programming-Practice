package main

import (
	"fmt"
)
func getRes(arr []int)int{
	m:=make(map[int]int)
	current:=arr[0]
	preSatis:=0
	res:=0
	for i:=1;i<len(arr);i++{
		if arr[i]==current{
			res+=preSatis
		}else{
			count:=0
			for k,_:=range m{
				if arr[i]>k{
					count++
				}
			}
			preSatis=count
			current=arr[i]
			res+=preSatis
		}
		m[arr[i]]=1

	}
	return res
}
func main(){
	var length int
	fmt.Scan(&length)
	var res []int
	for i:=0;i<length;i++{
		var tmp int
		fmt.Scan(&tmp)
		res=append(res,tmp)
	}
	fmt.Println(getRes(res))
}
