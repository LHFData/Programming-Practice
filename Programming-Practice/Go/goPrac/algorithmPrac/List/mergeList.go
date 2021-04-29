package main

type ListNode struct {
	Val int
	Next *ListNode
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var l1t *ListNode
	var l2t *ListNode
	l1t=l1
	l2t=l2


	temp:=new(ListNode)
	var tail *ListNode
	for l1t!=nil||l2t!=nil{
		res:=tail
		if l1t==nil&&l2t==nil{
			break
		}
		if l1t==nil&&l2t!=nil{
			res.Next=l2t
			l2t=nil
			break
		}else if l2t==nil&&l1t!=nil{
			res.Next=l1t
			l1t=nil
			break
		}else {

			if l1t.Val>l2t.Val{
				res.Next=l2t
				l2t=l2t.Next
				res=res.Next
				res.Next=nil
				tail=res
			}else{
				res.Next=l1t
				l1t=l1t.Next
				res=res.Next
				res.Next=nil
				tail=res
			}
		}
	}
	return temp.Next
}

func main(){
	n1:=[]int{1,3,4,5}
	n2:=[]int{2,3,6}
	l1:=new(ListNode)
	l2:=new(ListNode)
	var tal *ListNode
	tal=l1
	for i:=0;i<len(n1);i++{
		temp:=new(ListNode)
		temp.Val=n1[i]
		tal.Next=temp
		tal=tal.Next
	}

	tal=l2
	for i:=0;i<len(n2);i++{
		temp:=new(ListNode)
		temp.Val=n2[i]
		tal.Next=temp
		tal=tal.Next
	}
	mergeTwoLists(l1,l2)
}