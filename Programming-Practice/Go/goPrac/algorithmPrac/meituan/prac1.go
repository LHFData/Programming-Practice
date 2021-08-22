package main

import "fmt"

type way struct{
	srcx int
	srcy int
	disx int
	disy int
	cost int
}
func main(){
	var line,row,length int
	fmt.Scan(&line)
	fmt.Scan(&row)
	fmt.Scan(&length)

	var path [][]int

	for i:=0;i<length;i++{
		var temp []int
		for j:=0;j<5;j++{
			var num int
			fmt.Scan(&num)
			temp=append(temp,num)
		}
		path=append(path,temp)
	}
	getNext:=func(p [][]int,srcx int,srcy int) [][]int{
		var res [][]int
		for i:=0;i<len(p);i++{
			if (p[i][0]==srcx||p[i][0]==srcx+1)&&(p[i][1]==srcy||p[i][1]==srcy+1){
				res=append(res,p[i])
			}
		}
		return res
	}
	getMin:=func(pNext [][]int) []int{
		first:=pNext[0]
		for i:=0;i<len(pNext);i++{
			if pNext[i][4]<first[4]{
				first=pNext[i]
			}
		}
		return first
	}
	curx:=1
	cury:=1
	curcost:=0
	var Next [][]int
	Next=append(Next,[]int{1,1,1,1,0})
	for len(Next)>0{
		var temp [][]int
		temp=getNext(path,Next[len(Next)-1][2],Next[len(Next)-1][3])
		if len(temp)==0{
			Next=Next[:len(Next)-1]
			continue
		}
		if len(Next)==0&&curx!=line&&cury!=row{
			fmt.Println(-1)
			return
		}
		m:=getMin(temp)
		curcost=curcost+m[4]
		curx=m[2]
		cury=m[3]
		if curx==line&&cury==row{
			fmt.Println(curcost)
			return
		}
		Next=Next[:len(Next)-1]
		Next=append(Next,temp...)
	}




}
