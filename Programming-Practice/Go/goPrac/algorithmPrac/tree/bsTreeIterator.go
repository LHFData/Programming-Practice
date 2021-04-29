package main

import (
	"container/list"
	"fmt"
)

type BSTIterator struct {
	min int
	current int
	inorderseq []*TreeNode
    length int
}

func Constructor(root *TreeNode) BSTIterator {
	temp:=root
	var count int
	stack:=list.New()
	var inSeq []*TreeNode
	for temp.Left!=nil{
		temp=temp.Left
	}
	min:=temp.Val
	temp=root
	for temp!=nil||stack.Len()!=0{
		if temp!=nil{
			stack.PushBack(temp)
			temp=temp.Left
		}else{
			temp=stack.Remove(stack.Back()).(*TreeNode)
			inSeq=append(inSeq,temp)
			count++
			temp=temp.Right
		}
	}
	fmt.Println(inSeq)
	return BSTIterator{min: min,current: 0,inorderseq: inSeq,length:count}
}


func (this *BSTIterator) Next() int {

	this.current++
	return this.inorderseq[this.current].Val

}


func (this *BSTIterator) HasNext() bool {
	if this.current<this.length-1{
		return true
	}
	return false
}
func main(){
	t:=new(TreeNode)
	t.Val=1
	l:=new(TreeNode)
	l.Val=2
	r:=new(TreeNode)
	r.Val=0
	t.Left=r
	t.Right=l
	b:=Constructor(t)
	fmt.Println()
	fmt.Println(b.Next(),b.HasNext(),b.Next(),b.HasNext(),b.Next(),b.HasNext())
}
