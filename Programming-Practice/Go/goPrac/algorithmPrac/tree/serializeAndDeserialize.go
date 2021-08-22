package main

import (
	"fmt"
	"strconv"
	"strings"
)
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
type Codec struct {
	tree *TreeNode
	str string
}

func Constructor() Codec {
	return *new(Codec)
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root==nil{
		return ""
	}
	var temp []*TreeNode
	temp=append(temp,root)
	var res string
	res=res+strconv.Itoa(root.Val)+" "
	for len(temp)>0{
		var line []*TreeNode
		cur:=temp[0]
		if cur.Left!=nil{
			line=append(line,cur.Left)
			res=res+strconv.Itoa(cur.Left.Val)+" "
		}else{
			res=res+"null"+" "
		}
		if cur.Right!=nil{
			line=append(line,cur.Right)
			res=res+strconv.Itoa(cur.Right.Val)+" "
		}else{
			res=res+"null"+" "
		}
		temp=append(temp[1:], line...)
	}
	return res[:len(res)-1]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data)==0{
		return nil
	}
	strslice:=strings.Split(data," ")
	root:=new(TreeNode)
	//层数
	if strslice[0]=="null"{
		return nil
	}
	root.Val,_=strconv.Atoi(strslice[0])
	var temp []*TreeNode
	temp=append(temp,root)
	index:=0

	for len(temp)>0{
		var readin []*TreeNode
		supposeReadIn:=2*len(temp)
		for i:=1;i<=supposeReadIn;i++{
			if index+i>=len(strslice){
				break
			}
			if strslice[index+i]=="null"{
				readin=append(readin,nil)

			}else{
				t:=new(TreeNode)
				t.Val,_=strconv.Atoi(strslice[index+i])
				readin=append(readin,t)
			}
		}
		index=index+supposeReadIn
		if len(readin)==0{
			break
		}
		count:=0
		for i:=0;i<len(temp);i++{
			temp[i].Left=readin[count]
			count++
			temp[i].Right=readin[count]
			count++
		}
		for i:=0;i<len(readin);i++{
			if readin[i]==nil{
				readin=append(readin[:i], readin[i+1:]...)
				i--
			}
		}
		temp=readin
	}
	return root
}
func main(){
	c:=Constructor()
	r:=c.deserialize("4 -7 -3 null null -9 -3 9 -7 -4 null 6 null -6 -6 null null 0 6 5 null 9 null null -1 -4 null null null -2 null null null null null null null")
	fmt.Println(r)
}