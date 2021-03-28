package main

import (
	"fmt"
)

//type ListNode struct {
//	data int
//	next *ListNode
//}
func reverseList(head *ListNode) *ListNode {
	l:=head
	var prev *ListNode
	las:=l

	for las.next!=nil{
		temp:=las.next
		las.next=prev
		prev=las
		las=temp
	}
	las.next=prev
	return las
}


func main()  {
	var list *ListNode
	for i:=0;i<10;i++{
		l:=new(ListNode)
		l.data=i
		if list!=nil {
			l.next = list.next
			list.next = l
		}else{
			list=l
		}
	}
	list=reverseList(list)
	for list!=nil {
		fmt.Println(list)
		list=list.next
	}
}