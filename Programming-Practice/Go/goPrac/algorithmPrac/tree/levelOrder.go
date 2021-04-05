package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func levelOrder (root *TreeNode) [][]int{
	//r:=[]int{root.Val}
	var result [][]int
	//result=append(result,r)
	if root==nil{
		res:=[][]int{}
		return res
	}
	a:=[]*TreeNode{root}

	for len(a)!=0 {
		var level []*TreeNode
		var reslevel []int //输出结果
		for i := 0; i < len(a); i++ {
			reslevel=append(reslevel,a[i].Val)
			if a[i].Left != nil {
				//下一层遍历结果
				level = append(level, a[i].Left)
			}
			if a[i].Right != nil {
				//下一层遍历结果
				level = append(level, a[i].Right)
			}
		}
		a=level
		result = append(result, reslevel)
	}
	return result
}

func main(){
	var t,b,s,r,d TreeNode
	t.Val=10
	b.Val=11
	s.Val=2
	r.Val=3
	d.Val=4
	t.Right=&b
	t.Left=&s
	b.Left=&r
	s.Left=&d
	b.Right=nil
	s.Right=nil
	res:=levelOrder(&t)
	fmt.Println(res)

}