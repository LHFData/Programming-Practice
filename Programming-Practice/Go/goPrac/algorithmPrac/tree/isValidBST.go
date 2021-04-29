package main
import "container/list"
func isValidBST(root *TreeNode) bool {
	stack:=list.New()
	temp:=root
	var res []int
	stack.PushBack(temp)
	for stack.Len()>0||temp!=nil{
		for temp!=nil{
			stack.PushBack(temp)
			temp=temp.Left
		}
		for stack.Len()>0{
			temp=stack.Remove(stack.Back()).(*TreeNode)
			res=append(res,temp.Val )
			temp=temp.Right
		}
	}
	for i:=0;i+1<len(res);i++{
		if res[i]>=res[i+1]{
			return false
		}
	}
	return true
}
