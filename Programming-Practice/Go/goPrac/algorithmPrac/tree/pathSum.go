package main

import "fmt"
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func pathSum(root *TreeNode, target int) [][]int {
	var res [][]int
	var pre func(node *TreeNode,count int,path []int)
	pre=func (node *TreeNode,count int,path []int){
		if node.Right==nil&&node.Left==nil&&count-node.Val==0{
			path=append(path,node.Val)
			res=append(res,path)
		}
		if node.Right==nil&&node.Left==nil{
			return
		}else{
			if node.Left!=nil{
				left:=append(path,node.Val)
				pre(node.Left,count-node.Val,left)
			}
			if node.Right!=nil{
				right:=append(path,node.Val)
				pre(node.Right,count-node.Val,right)
			}

		}

	}

	pre(root,target,[]int{})
	return res

}
func TreeBuild(t []int) *TreeNode{

	arr:=make([]*TreeNode,len(t))
	i:=0
	for i<len(t){
		if t[i]!=-1{
			temp:=new(TreeNode)
			temp.Val=t[i]
			arr[i]=temp
		}
		i++
	}
	for i=0;i<len(arr);i++{
		if arr[i]!=nil{
			if i*2<len(arr){
				if arr[i*2]!=nil{
					arr[i].Left=arr[i*2]
				}
			}
			if i*2+1<len(arr){
				if arr[i*2+1]!=nil{
					arr[i].Right=arr[i*2+1]
				}
			}
		}
	}
	return arr[0]
}
func main(){
	t:=TreeBuild([]int{5,4,8,11,-1,13,4,7,2,5,1})
	fmt.Println(pathSum(t,3))
}