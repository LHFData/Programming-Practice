package main

func hasDFS(root *TreeNode,pre int,target int) bool{
	if root==nil{
		return false
	}
	if root.Right==nil&&root.Left==nil{
		if pre+root.Val==target{
			return true
		}
		return false
	}

	LeftRe:=hasDFS(root.Left,pre+root.Val,target)
	RightRe:=hasDFS(root.Right, pre+root.Val, target)
	if LeftRe||RightRe{
		return true
	}
	return false

}
func hasPathSum(root *TreeNode, targetSum int) bool {
	return hasDFS(root,0,targetSum)
}
