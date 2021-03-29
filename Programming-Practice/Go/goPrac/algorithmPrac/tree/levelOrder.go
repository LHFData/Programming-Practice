package main


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func levelOrder (root *TreeNode) [][]int{
	r:=[]int{root.Val}
	var result [][]int
	result=append(result,r)
	a:=[]*TreeNode{root}
	for len(a)!=0 {
		var level []*TreeNode
		var reslevel []int
		for i := 0; i < len(a); i++ {
			reslevel=append(reslevel,a[i].Val)
			if a[i].Left != nil {
				level = append(level, a[i].Left)
			}
			if a[i].Right != nil {
				level = append(level, a[i].Right)
			}
		}
		result = append(result, reslevel)
	}
	return result
}
