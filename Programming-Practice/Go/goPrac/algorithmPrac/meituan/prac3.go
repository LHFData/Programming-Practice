package main

import "fmt"

func main(){
	var nums int
	fmt.Scan(nums)
	//var curstat,profit,damage,totalTime int
	var input [][4]int
	for i:=0;i<nums;i++{
		var temp [4]int
		for j:=0;j<4;j++{
			var inc int
			fmt.Scan(&inc)
			temp[j]=inc
		}
		input=append(input,temp)
	}
	for i:=0;i<nums;i++{
		t:=input[i][3]
		curstat:=
		var dp [][2]int
		for j:=0;j<t;j++{
			var temp [2]int
			if j==0{
				if input[i][0]<input[i][1]{

				}
			}
		}
	}
}