package main

import (
	"fmt"
)
func isSeg(s []byte)bool{
	if len(s)==1{
		return true
	}
	var k int
	k=1
	for k<=len(s)/2{
		if equal(s[:k],s[len(s)-k:]){
			return true
		}else{
			k++
		}
	}
	return false
}
func equal(a,b []byte)bool{
	if len(a)!=len(b){
		return false
	}
	for i:=0;i<len(a);i++{
		if a[i]!=b[i]{
			return false
		}
	}
	return true
}

func backTrack(s []byte,Res *int){

	if len(s)<=1{
		*Res++
		return
	}
	for i:=1;i<=len(s);i++{
		if isSeg(s[0:i]){
			backTrack(s[i:],Res)
		}
	}
}
func solution(s []byte)int{
	n:=len(s)
	res:=0
	var backtrack func(idx int)
	backtrack=func(idx int){
		if idx>=n{
			res++
			return
		}
		for i:=1;idx+i<n;i++{
			if isSeg(s[idx:idx+i]){
				backtrack(idx+i)
			}
		}
	}
	backtrack(0)
	return res
}
func main(){
	var input string
	fmt.Scan(&input)
	s:=[]byte(input)
	res:=0
	backTrack(s,&res)
	fmt.Println(res)
}
