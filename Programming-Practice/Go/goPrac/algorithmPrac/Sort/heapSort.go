package main

import "fmt"

func main(){
	datas:=[]int{
		4,1,5,7,8,1,2,12,63,64,23,46,76,7,123,12535,3,
	}
	heapSort(datas)
	fmt.Println(datas)
}
func heapSort(s []int){
	n:=len(s)-1
	for k:=n/2;k>=1;k--{
		sink(s,k,n)
	}
	for n>1{
		swap(s,1,n)
		n--
		sink(s,1,n)
	}
}
func sink(s []int,k,n int)  {
	for {
		i:=2*k
		if i>n{
			break
		}
		if i<n&&s[i+1]>s[i]{
			i++
		}
		if s[k]>s[i]{
			break
		}
		swap(s,k,i)
		k=i
	}
}
func swap(s []int,i int,j int){
	s[i],s[j]=s[j],s[i]
}