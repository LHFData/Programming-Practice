package output

import (
	"crypto/rand"
	"fmt"
	"math/big"
	//"time"
)
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
func(l *ListNode) PrintList(){
	for l!=nil{
		fmt.Print(l.Val,"->")
		l=l.Next
	}
	fmt.Println()
}
func(l *ListNode) ArrayToListInHeadOrder(array []int) *ListNode{
	var head=new(ListNode)
	var list *ListNode
	list=head
	for i:=0;i<len(array);i++{
		val:=array[i]
		list.Next=&ListNode{val,list.Next}
	}
	return head.Next
}
//没有问题
func(l *ListNode) ArrayToListInTailOrder(array []int) *ListNode{
	var head=new(ListNode)
	var list,tail *ListNode
	list=head
	tail=head
	for i:=0;i<len(array);i++{

		val:=array[i]
		tail.Next=&ListNode{val,nil}
		tail=tail.Next
	}

	return list.Next
}
func(*ListNode) OutputRandomList(len int,num int,rangeofNum int64) []*ListNode{
	var listArray =make([]*ListNode,num+1)
	for j:=0;j<=num;j++ {
		var head =new(ListNode)
		var list *ListNode
		list=head
		for i := len; i >= 0; i-- {
			//rand.Seed(time.Now().Unix())
			val,_:=rand.Int(rand.Reader,big.NewInt(rangeofNum))
			list.Next=&ListNode{int(val.Int64()),list.Next}
		}
		listArray[j]=head.Next
	}
	return listArray
}

/*func main()  {
	var l ListNode
	listarray:=l.OutputRandomList(2,3)
	for i:=0;i<len(listarray);i++ {
		fmt.Println("list:")
		listarray[i].PrintList()
	}
}*/