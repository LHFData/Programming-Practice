package main
 type TreeNode struct {
    Val int
     Left *TreeNode
     Right *TreeNode
 }
func pre(r *TreeNode,res []int){
	if r!=nil{
		pre(r.Left,res)
		pre(r.Right,res)
		res=append(res,r.Val)
	}else{
		return
	}
}
func preorderTraversal(root *TreeNode) []int {
	var res []int
	pre(root,res)
	return res

}
