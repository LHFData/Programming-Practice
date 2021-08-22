package main

import "fmt"

func pathSum(root *TreeNode, target int) [][]int {
	var res [][]int
	var pre func(node *TreeNode,count int,path []int)
	pre=func (node *TreeNode,count int,path []int){
		if node.Right==nil&&node.Left==nil&&count-node.Val==0{
			path=append(path,node.Val)
			temp:=make([]int,len(path))
			copy(temp,path)
			res=append(res,temp)
			path=path[:len(path)-1]
		}
		if node.Right==nil&&node.Left==nil{
			return
		}
		if node.Left!=nil{
			left:=append(path,node.Val)
			pre(node.Left,count-node.Val,left)
		}
		if node.Right!=nil{
			right:=append(path,node.Val)
			pre(node.Right,count-node.Val,right)
		}
	}



	pre(root,target,[]int{})
	return res

}
func TreeBuild(t []int) *TreeNode{

	arr:=make([]*TreeNode,len(t))
	i:=0
	for i<len(t){
		if t[i]!=-1{
			temp:=new(TreeNode)
			temp.Val=t[i]
			arr[i]=temp
		}
		i++
	}
	for i=0;i<len(arr);i++{

		if arr[i]!=nil{
			if 2*i+1<len(arr){
					arr[i].Left=arr[2*i+1]

			}
			if i*2+1<len(arr){

					arr[i].Right=arr[(i+1)*2]

			}
		}

	}
	return arr[0]
}
func main(){
	t:=TreeBuild([]int{5,4,8,11,-1,13,4,7,2,-1,-1,-1,-1,5,1})
	fmt.Println(pathSum(t,22))
}