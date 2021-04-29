package main

import (

	"container/list"
	"fmt"
)

//func Pre(t *TreeNode,res *[]int){
//	if t!=nil {
//		*res = append(*res, t.Val)
//		Pre(t.Left,res)
//		Pre(t.Right,res)
//	}
//}
//func preorderTraversal(root *TreeNode) []int {
//	var res []int
//	Pre(root,&res)
//	return res
//}
func preorderTraversal(root *TreeNode) []int {
	var res []int
	stack:=list.New()
	if root==nil{
		return nil
	}
	stack.PushBack(root)

	for stack.Len()>0{
		temp:=stack.Remove(stack.Back()).(*TreeNode)
		res=append(res,temp.Val)
		if temp.Right!=nil{
			stack.PushBack(temp.Right)
		}
		if temp.Left!=nil{
			stack.PushBack(temp.Left)
		}

	}
	return res
}

func main(){
	t:=new(TreeNode)
	t.Val=1
	l:=new(TreeNode)
	l.Val=2
	r:=new(TreeNode)
	r.Val=0
	t.Left=l
	t.Right=r
	fmt.Println(preorderTraversal(t))
}