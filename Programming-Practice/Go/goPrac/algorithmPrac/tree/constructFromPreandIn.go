package main

func getRoot(preorder []int,inorder []int)*TreeNode{
	if len(preorder)==0{
		return nil
	}
	root:=preorder[0]
	var res *TreeNode
	partition:=func(in []int,r int)(left []int,right []int,leftroot int){
		var temp int
		temp=-1
		for i:=0;i<len(in);i++{
			if in[i]==r{
				temp=i
			}
		}

		return in[:temp],in[temp+1:],in[temp]


	}
	getNextPreorder:=func(pre []int,leftroot int)([]int,[]int){
		var temp int
		for i:=0;i<len(pre);i++{
			if pre[i]==leftroot{
				temp=i
			}
		}
		return pre[1:temp+1],pre[temp+1:]
	}
	left,right,leftroot:=partition(preorder,root)//中序划分
	preLeft,preRight:=getNextPreorder(preorder,leftroot)//根据中序的前一个值找到前序的左右子树，得到左右子树的前序
	res.Left=getRoot(preRight, left)
	res.Right=getRoot(preLeft, right)//
	res.Val=root
	return res
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return getRoot(preorder,inorder)
}
