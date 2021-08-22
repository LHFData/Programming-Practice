package main

import "fmt"

func main(){
	var wallNum,width,maxHeight int
	fmt.Scan(&wallNum)
	fmt.Scan(&width)
	fmt.Scan(&maxHeight)
	left:=0
	right:=0
	var wall []int
	for i:=0;i<wallNum;i++{
		var temp int
		fmt.Scan(&temp)
		wall=append(wall,temp)
	}
	for right<wallNum{
		if right-left>=width&&right+1<wallNum&&wall[right+1]<=maxHeight{
			right++
			continue
		}
		if right==left&&wall[right]>maxHeight{
			left++
			right++
		}else if right!=left&&wall[right]>maxHeight&&right-left<width{
			right++
			left=right
		}else if right!=left&&wall[right]>maxHeight&&right-left>=width{
			break
		}else if wall[right]<=maxHeight{
			right++
		}

	}
	if right-left>width{
		mini:=left
		for i:=left;i<right-width;i++{
			if wall[i]<wall[mini]{
				mini=i
			}
		}
		fmt.Println(mini)
		return
	}else if right-left==width{
		fmt.Println(left)
	}else{
		fmt.Println(-1)
	}
}
