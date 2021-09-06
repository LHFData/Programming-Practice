package main
type ListNode struct{
	Next *ListNode
	Val int
}
func reverseKGroup( head *ListNode ,  k int ) *ListNode {
	temp:=head

	//翻转
	reverse:=func(subHead *ListNode)*ListNode{
		var pre *ListNode
		tmp:=subHead
		for tmp!=nil{
			next:=tmp.Next
			tmp.Next=pre
			tmp=next
		}
		return pre
	}
	count:=0
	var resList *ListNode
	resList=temp
	var newHead,res *ListNode
	//找处理部分
	flag:=true
	for temp!=nil{
		if count==k{
			tmp:=temp.Next
			temp.Next=nil
			newHead=reverse(resList)
			if flag&&resList==temp{
				flag=false
				res=newHead
			}
			resList.Next=tmp
			resList=tmp
			count=0
		}else{
			temp=temp.Next
			count++
		}
	}

	return res

}
