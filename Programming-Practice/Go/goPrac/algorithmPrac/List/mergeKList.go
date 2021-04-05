package List

import (
	"fmt"
)
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
func PrintList(node *ListNode){
	temp:=node
	for temp!=nil{
		fmt.Println(temp.Val)
		temp=temp.Next
	}
}
func mergeKLists(lists []*ListNode) *ListNode {
	length:=len(lists)

	var res *ListNode
	//指针声明之后一定要初始化
	var restail *ListNode

	getMin:= func(l []*ListNode) (int){
		var index int
		var val int
		l=lists
		count:=0
		for i:=0;i<len(l);i++{
			if l[i]!=nil{
				index=i
				val=l[i].Val
				break
			}
		}

		for i:=0;i<length;i++{
			if l[i]==nil{
				count++
				continue
			}
			if l[i].Val<val{
				val=l[i].Val
				index=i
			}
		}

		return index
	}
	Empty:=func(l []*ListNode) bool{
		for i:=0;i<length;i++{
			if l[i]!=nil{
				return false
			}
		}
		return true
	}
	for !Empty(lists){
		i:=getMin(lists)
		if res==nil{
			res=lists[i]
			restail=lists[i]
			lists[i]=lists[i].Next
		}else{
			restail.Next=lists[i]
			restail = restail.Next
			if lists[i].Next!=nil {
				lists[i] = lists[i].Next
			}else {
				lists[i]=nil
			}

		}

	}
	return res
}

func main(){
	test:=CreateList([][]int{{0},{1}})
	PrintList(mergeKLists(test))
	//fmt.Println(mergeKLists(test))

}