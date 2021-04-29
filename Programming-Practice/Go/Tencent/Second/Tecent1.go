

package main
import . "nc_tools"

  type ListNode struct {
	  Val  int
	  Next *ListNode
  }

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param S ListNode类 val表示权值，next指向下一个元素
 * @return ListNode类
 */

func solve( S *ListNode ) *ListNode {
	var minGroup []*ListNode
	var minVal int
	var end *ListNode
	temp:=S
	minVal=S.Val
	for temp!=nil{
		if temp.Val<minVal{
			minVal=temp.Val
		}
		temp=temp.Next
		if temp.Next==nil{
			end=temp
		}
	}
	temp=S
	for temp!=nil{
		if temp.Val==minVal{
			minGroup=append(minGroup,temp)
		}
	}
	end.Next=S
	for i:=0;i<len(minGroup);i++{
		temp=minGroup[i].Next
		var next []int

		for temp!=minGroup[i]{


		}
	}
	// write code here
}
