package main

import "fmt"

func getRes4(arr [][3]int,grass [][]int,m int,n int) int{
	var res int

	for d:=0;d<len(arr);d++{
		for i:=0;i<m;i++{
			for j:=0;j<n;j++{
				if (i - arr[d][0])*(i - arr[d][0]) + (j - arr[d][1])*(j - arr[d][1])<= (arr[d][2])*(arr[d][2]){
					grass[i][j]=0
				}else{
					grass[i][j]++
				}
			}
		}
	}
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			res+=grass[i][j]
		}
	}

	return res
}
func main(){
	var m,n int
	var days int
	fmt.Scan(&m)
	fmt.Scan(&n)
	fmt.Scan(&days)
	var input [][3]int

	grass:=make([][]int,m)
	for i:=0;i<m;i++{
		temp:=make([]int,n)
		grass=append(grass,temp)
	}
	for i:=0;i<days;i++{
		tmp1,tmp2,tmp3:=0,0,0
		fmt.Scan(&tmp1)
		fmt.Scan(&tmp2)
		fmt.Scan(&tmp3)
		var tmp [3]int
		tmp[0]=tmp1
		tmp[1]=tmp2
		tmp[2]=tmp3
		input=append(input,tmp)
	}

	fmt.Println(getRes4(input,grass,m,n))
}
