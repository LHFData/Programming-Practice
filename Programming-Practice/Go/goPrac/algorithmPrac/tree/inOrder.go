package main

import (
	"container/list"
	"fmt"
)

func inorderTraversal(root *TreeNode) []int {
	stack:=list.New()
	var res []int
	if root==nil{
		return nil
	}
	var temp *TreeNode

	temp=root
	for temp!=nil||stack.Len()!=0{
		if temp!=nil{
			stack.PushBack(temp)
			temp=temp.Left

		}else{
			temp=stack.Remove(stack.Back()).(*TreeNode)
			res=append(res,temp.Val)
			temp=temp.Right
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
	fmt.Println(inorderTraversal(t))
}