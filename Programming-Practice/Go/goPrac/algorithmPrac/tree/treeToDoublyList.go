package main

import "fmt"

type Node struct {
	 Val int
	 Left *Node
	 Right *Node
}
func treeToDoublyList(root *Node) *Node{
	var stack []*Node
	if root==nil{
		return nil
	}
	var p *Node
	var head *Node
	var flag bool
	var res []*Node
	p=root
	for p!=nil||len(stack)>0{
		if p!=nil{
			stack=append(stack, p)
			if p.Left==nil&&!flag{
				head=p
				flag=true
			}
			p=p.Left

		}else{
			p=stack[len(stack)-1]
			if len(res)==0{

				res=append(res,p)
			}else{
				res[len(res)-1].Right=p
				p.Left=res[len(res)-1]
				res=append(res,p)
			}
			stack=stack[:len(stack)-1]
			p=p.Right
		}
	}
	res[len(res)-1].Right=res[0]
	res[0].Left=res[len(res)-1]
	return head
}
func NodeTreeBuild(t []int) *Node{

	arr:=make([]*Node,len(t))
	i:=0
	for i<len(t){
		if t[i]!=-1{
			temp:=new(Node)
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
func main (){
	t:=treeToDoublyList(NodeTreeBuild([]int{4,2,5,1,3}))
	fmt.Println(t)
}