package main

import "fmt"

type TreeNode struct {
 Val int
 Left *TreeNode
 Right *TreeNode
 }
func solve( root *TreeNode ) *TreeNode {

	var a []*TreeNode
	a=append(a,root)
	flag:=false
	for a!=nil{
		var temp []*TreeNode
		for i:=0;i<len(a);i++{
			if flag{
				a[i].Left=nil
				a[i].Right=nil
			}
			if a[i].Left!=nil{
				temp=append(temp,a[i].Left)
			}else{
				flag=true
			}
			if a[i].Right!=nil{
				temp=append(temp,a[i].Right)
			}
		}
		if len(temp)==0{

		}
		a=temp
	}
	return root
	// write code here
}

func main(){
	h:=&TreeNode{Val:8,Left: nil,Right:nil}
	//g:=&TreeNode{Val:7,Left: h,Right:nil}
	//f:=&TreeNode{Val:6,Left: nil,Right:nil}
	a:=&TreeNode{Val:5,Left: h,Right:nil}
	//b:=&TreeNode{Val:4,Left: nil,Right:nil}
	c:=&TreeNode{Val:3,Left: nil,Right:nil}
	d:=&TreeNode{Val:2,Left: nil,Right:a}
	e:=&TreeNode{Val:1,Left: d,Right:c}
	res:=solve(e)
	fmt.Println(res)
}