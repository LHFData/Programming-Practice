package main

import "fmt"

func insert(a []int,b[]int)[]int{
	var temp []int
	i,j:=0,0
	for i<len(a)&&j<len(b){
		if less(a[i],b[j]){
			temp=append(temp,a[i])
			i++
		}else{
			temp=append(temp,b[j])
			j++
		}
	}
		temp=append(temp,a[i:]...)


		temp=append(temp,b[j:]...)

	return temp
}
func less(a,b int)bool{
	if a<b{
		return true
	}else{
		return false
	}
}
func MergeSort(arr []int) []int{
	if len(arr)<2{
		return arr
	}
	mid:=len(arr)/2
	left:=MergeSort(arr[:mid])
	right:=MergeSort(arr[mid:])
	return insert(left,right)
}

func main(){
	a:=[]int{3,1,4,2,5,7,6}
	aa:=MergeSort(a)
	fmt.Println(aa)
}