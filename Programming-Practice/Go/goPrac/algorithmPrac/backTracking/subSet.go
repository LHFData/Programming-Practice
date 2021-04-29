package main

import "fmt"

func DFS(res *[][]int,nums []int,visit []int,level int,choose bool) {
	if level==len(nums)-1{
		if choose{
			visit=append(visit,nums[level])
		}
		*res=append(*res,visit)
		return
	}
	if choose{
		visit=append(visit,nums[level])
	}
	left:=make([]int,len(visit))
	copy(left,visit)
	right:=make([]int,len(visit))
	copy(right,visit)
	DFS(res,nums,left,level+1,true)
	DFS(res,nums,right,level+1,false)
	//数组和切片使用的方式区别
	//只有完全拷贝才能够避免参数传递返回时对原数组的影响

}
//func ex(a []int,b []int){
//	temp:=a[0]
//	a[0]=b[0]
//	b[0]=temp
//}
func subsets(nums []int) [][]int {
	var res [][]int
	var visit []int
	DFS(&res,nums,visit,0,true)
	DFS(&res,nums,visit,0,false)
	return res
}

func main()  {
	//a:=[]int{1,2}
	//b:=[]int{3,4}
	//c:=make([]int,len(b))
	//copy(c,b)
	//c[0]=6
	//ex(a,b)
	//fmt.Println(c,b)
	fmt.Println(subsets([]int{1,2,3,4,5,6,7}))
}
