package main

import (
	"container/list"
)


func countNodes(root *TreeNode) int {
	stack:=list.New()
	count:=0
	if root==nil{
		return 0
	}
	stack.PushBack(root)
	for stack.Len()!=0{
		temp:=stack.Remove(stack.Back()).(*TreeNode)
		count++
		if temp.Right!=nil{
			stack.PushBack(temp.Right)
		}
		if temp.Left!=nil{
			stack.PushBack(temp.Left)
		}
	}
	return count
}
