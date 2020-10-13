package main

import "fmt"

func isParlindrome(x int) bool{
	xTemp:=x
	array:=[]int{}
	count:=0
	if xTemp<0{
		return false
	}
	for xTemp>=1{
		array=append(array,xTemp%10)
		xTemp=xTemp/10
		count++
	}
	fmt.Println(array,count)
	for left,right:=0,0;left+right<count;{
		if array[left]==array[count-1-right]{
			left++
			right++
		}else {
			return false
		}
	}
	return true

}
func main(){
	fmt.Println(isParlindrome(21741247))
}
