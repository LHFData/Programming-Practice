package main

import (
	"container/heap"
	"fmt"
	)
type arr []int
func(a arr)Len()int{return len(a)}
func(a arr)Less(i int,j int)bool{return a[i]>a[j]}
func(a arr)Swap(i int,j int){a[i],a[j]=a[j],a[i]}
func(a *arr)Push(x interface{}){
	*a=append(*a,x.(int))
}
func(a *arr)Pop() interface{}{
	old:=*a
	n:=len(old)
	x:=old[n-1]
	*a=old[0:n-1]
	return x
}

func main(){
	h:=&arr{3,1,2,4,5}
	heap.Init(h)
	heap.Push(h,6)
	for _,ele:=range *h{
		fmt.Println("%d",ele)
	}
	fmt.Println("after push")
	for i,ele:=range *h{
		if ele==3{
			(*h)[i]=10
			heap.Fix(h,i)
		}
	}
	for _,ele:=range *h{
		fmt.Println("%d",ele)
	}
	fmt.Println("After Fix")
	lhf:=heap.Pop(h)
	fmt.Println("%d",lhf)
	for _,ele:=range *h{
		fmt.Println("%d",ele)
	}
}