package main

type ListNode struct {
	Val int
	Next *ListNode
}
func CreateList(nums [][]int) []*ListNode{
	var res []*ListNode
	for i:=0;i<len(nums);i++{
		var head *ListNode
		tailForInsert:=new(ListNode)
		for j:=0;j<len(nums[i]);j++{
			temp:=new(ListNode)
			temp.Val=nums[i][j]
			if head==nil{
				head=temp
				tailForInsert=temp
			}else{
				tailForInsert.Next=temp
				tailForInsert=tailForInsert.Next
			}
		}
		res=append(res,head)
	}
	return res
}
func hasCycle(head *ListNode) bool {
	var fast *ListNode
	var slow *ListNode
	nilHead:=new(ListNode)
	nilHead.Next=head
	count:=0
	fast=nilHead
	slow=nilHead
	if head==nil{
		return false
	}else if head.Next==nil{
		return false
	}
	for {

		fast=fast.Next
		fast=fast.Next
		slow=slow.Next
		if fast.Next==nil{
			return false
		}
		if fast==slow{
			break
		}
	}
	slow=head
	for slow!=fast{
		count++
		slow=slow.Next
	}


	return true
}

func main()  {
	input:=CreateList([][]int{{1,3,4,5,6,7},{1,4,5,3,3,2,7,8}})
	head:=input[1]
	tail:=input[1]
	for i:=0;i<8;i++{
		if i<3 {
			head = head.Next
		}
		if tail.Next!=nil{
			tail=tail.Next
		}
	}
	tail.Next=head
	hasCycle(input[1])
}