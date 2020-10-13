package main


import (
	"fmt"
	"math"
	"./output"
)
/*
type ListNode struct {
	Val  int
	Next *ListNode
}*/

func addTwoNumbersSmallInt(l1 *output.ListNode,l2 *output.ListNode) *output.ListNode{
	var num1 int
	var num2 int
	num1=0
	num2=0
	for e,i:=l1,0;e!=nil;e,i=e.Next,i+1{
		num1=num1+e.Val*int(math.Pow(10,float64(i)))
		//fmt.Println("we get num1 =",num1)
	}
	for f,i:=l2,0;f!=nil;f,i=f.Next,i+1{
		num2=num2+f.Val*int(math.Pow(10,float64(i)))
		//fmt.Println("we get num2 =",num2)
	}
	outputnum:=num1+num2
	fmt.Println(outputnum)
	var list,tail *output.ListNode
	head:=new(output.ListNode)
	list=head
	tail=head
	for {
		tail.Next=&output.ListNode{outputnum%10,nil}
		tail=tail.Next
		outputnum=outputnum/10
		if outputnum<1 {
			break
		}

	}

	fmt.Println(num1,num2)
	list.Next.PrintList()
	return l1
}

func addTwoNumbers(l1 *output.ListNode,l2 *output.ListNode) *output.ListNode{
	var ten bool
	ten=false
	var Output,list *output.ListNode
	head:=new(output.ListNode)
	Output=head
	list=head
	for i:=0;l1!=nil||l2!=nil;i++{
		var sum int
		if l1==nil&&l2!=nil{
			if ten{
				sum=l2.Val+1
			} else {
				sum=l2.Val
			}
		} else if l1!=nil&&l2==nil{
			if ten{
				sum=l1.Val+1
			}else{
				sum=l1.Val
			}
		} else{
			if ten{
				sum=l1.Val+l2.Val+1
			}else{
				sum=l1.Val+l2.Val
			}
		}
		if sum>9{
			sum=sum-10
			ten=true
		} else{
			ten=false
		}

		Output.Next=&output.ListNode{sum,nil}
		Output=Output.Next
		if l1!=nil{
			if l1.Next==nil{
				l1=nil
			}else{
				l1=l1.Next
			}
		}
		if l2!=nil{
			if l2.Next==nil {
				l2 = nil
			}else{
				l2=l2.Next
			}
		}
	}
	if ten{
		Output.Next=&output.ListNode{1,nil}
		Output=Output.Next
	}
	list.Next.PrintList()
	return list.Next
}
func main(){
	var list output.ListNode
	//input:=list.OutputRandomList(3,2,10)
	//input[0].PrintList()
	//input[1].PrintList()
	t1:=list.ArrayToListInTailOrder([]int{1,2})
	t1.PrintList()
	t2:=list.ArrayToListInTailOrder([]int{9,9,9})
	addTwoNumbers(t1,t2)
	//addTwoNumbers(input[0],input[1])

}

