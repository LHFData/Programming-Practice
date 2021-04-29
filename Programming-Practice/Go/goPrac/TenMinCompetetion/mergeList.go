package main

type ListNode struct{
	Val int
	Next *ListNode
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode
	for l1!=nil||l2!=nil{
		if l1==nil{
			l3.Next=l2
			break
		}else if l2==nil{
			l3.Next=l1
			break
		}
		if l1.Val<l2.Val{
			l3.Next=l1
			l3=l3.Next
			l1=l1.Next
		}else {
			l3.Next=l2
			l3=l3.Next
			l2=l2.Next
		}
	}
	return l3
}

func main(){

}
