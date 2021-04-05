package main

import "fmt"

//type ListNode struct {
//    Val int
//    Next *ListNode
//}
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	var tail *ListNode
	var prev *ListNode
	prev=new(ListNode)
	prev.Next=head
	count:=0
	front:=prev
	tail=head
	for tail!=nil{
		if count<n{
			count++
			tail=tail.Next
		}else{
			front=front.Next
			tail=tail.Next
		}
	}
	if front==prev{
		h:=head.Next
		head.Next=nil
		return h
	}
	if front.Next==nil{
		head=nil
		return head
	}
	if front.Next.Next!=nil{
		front.Next=front.Next.Next
	}else{
		front.Next=nil
	}
	return head
}
func main()  {
	var list *ListNode
	for i:=0;i<5;i++{
		l:=new(ListNode)
		l.Val=i
		if list!=nil {
			l.Next = list.Next
			list.Next = l
		}else{
			list=l
		}
	}
	removeNthFromEnd(list,1)
	fmt.Println(list.Val)

}
