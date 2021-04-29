package main

import "fmt"

func LevelOrder(root *TreeNode) [][]int {
	var res [][]int
	var temp []*TreeNode
	flag:=true
	temp=append(temp, root)
	if root==nil{
		return nil
	}
	for len(temp)>0{
		var newtemp []*TreeNode
		var inttemp []int

		for i:=0;i<len(temp);i++{
			if temp[i].Left!=nil{
				newtemp=append(newtemp,temp[i].Left)
				}
			if temp[i].Right!=nil{
				newtemp=append(newtemp, temp[i].Right)
			}
		}
		if flag{
			for i:=0;i<len(temp);i++{
				inttemp=append(inttemp,temp[i].Val)
			}
			flag=false
		}else{
			for i:=len(temp)-1;i>=0;i--{
				inttemp=append(inttemp,temp[i].Val)
			}
			flag=true
		}
		res=append(res, inttemp)
		temp=newtemp

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
	fmt.Println(LevelOrder(t))
}