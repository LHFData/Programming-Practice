package main

import "container/list"

func postOrder(root *TreeNode,target int)[][]int{
	stack:=list.New()
	var res [][]int
	var pre *TreeNode
	stack.PushBack(stack)
	for stack.Len()>0{
		temp:=stack.Back().Value.(*TreeNode)
		if (temp.Left==nil&&temp.Right==nil)||(pre!=nil&&(pre==temp.Right||pre==temp.Left)){
			if temp.Left==nil&&temp.Right==nil{
				var count int
				var p []int
				for e:=stack.Back();e!=nil;e=e.Next(){
					count=count+e.Value.(*TreeNode).Val
					p=append(p,e.Value.(*TreeNode).Val)
				}
				if count+temp.Val==target{
					p=append(p,temp.Val)
					res=append(res,p)
				}
			}
			pre=temp
			stack.Remove(stack.Back())
		}else{
			if temp.Right!=nil{
				stack.PushBack(temp.Right)
			}
			if temp.Left!=nil{
				stack.PushBack(temp.Left)
			}
		}
	}
	return res
}
