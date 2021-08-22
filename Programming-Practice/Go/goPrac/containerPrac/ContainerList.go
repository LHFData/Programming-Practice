package main
import ("container/list"
	"fmt"
)

func print(l *list.List){
	for e:=l.Front();e!=nil;e=e.Next(){
		fmt.Println(e.Value)
	}
}
type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}
func main(){
	//l:=list.New()
	//s:=&TreeNode{2,nil,nil}
	//l.PushBack(s)
	//ss:=l.Back().Value.(*TreeNode)
	//fmt.Println(ss.Val)
	//
	s:=3%3
	fmt.Println(s)
}
