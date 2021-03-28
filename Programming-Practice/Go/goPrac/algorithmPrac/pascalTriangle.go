package main

import "fmt"

func pascal(k int) [][]int{
	var result [][]int

	result=append(result,[]int{1})
	for i:=1;i<k;i++{
		var resl []int
		for l:=0;l<i+1;l++{
			var left int
			var right int
			if l-1<0{
				left=0
			} else {
				left=result[i-1][l-1]
			}
			if l>=i{
				right=0
			}else {
				right = result[i-1][l]
			}

			resl=append(resl,left+right)
		}
		result=append(result,resl)
	}
	return result
}
func main(){
	fmt.Println(pascal(5))
}