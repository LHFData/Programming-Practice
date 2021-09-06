package main
import (
	"fmt"
)

func main(){
	var n,l,r int
	fmt.Scan(&n)
	fmt.Scan(&l)
	fmt.Scan(&r)
	var res []int

	res=append(res,n)
	for res[0]!=1||res[len(res)-1]!=1{
		var temp []int
		for i:=0;i<len(res);i++{
			if res[i]==1||res[i]==0{
				temp=append(temp,res[i])
				continue
			}
			a,b:=res[i]/2,res[i]%2
			temp=append(temp,a)
			temp=append(temp,b)
			temp=append(temp,a)
		}
		res=temp
	}
	count:=0
	for i:=l-1;i<=r-1;i++{
		if res[i]==1{
			count++
		}
	}
	fmt.Println(count)
}
