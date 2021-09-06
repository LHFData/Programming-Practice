package main

import (
	"container/heap"
	"fmt"
	"sort"
)
type arr []int

func (a arr ) Less(i,j int)bool{
	return a[i]<a[j]
}
func (a *arr) Push(n interface{}){
	*a=append(*a,n.(int))
}
func (a *arr) Pop()interface{}{
	old:=*a
	n:=len(old)
	tmp:=old[n-1]
	*a=old[:n-1]
	return tmp
}
func (a arr) Swap(i,j int){
	a[i],a[j]=a[j],a[i]
}
func (a arr) Len()int{
	return len(a)
}
type list []int
func main(){
	arry:=[][]int{{1,3,4},{3,2,1},{4,1,2}}
	sort.Slice(arry, func(i, j int) bool {
		return arry[i][1]<arry[j][1]
	})
	fmt.Println(arry)
	a:=arr{1,3,6,2,7}
	var aa list
	aa=list(a)
	heap.Push(&a,4)
	heap.Init(&a)
	fmt.Println(a)
	fmt.Println(heap.Pop(&a))
	fmt.Println(heap.Pop(&a))
	fmt.Println(heap.Pop(&a))
	fmt.Println(heap.Pop(&a))
}
