package main

import (
	"fmt"
)

func sort(arr []int,start int ,end int){
	if start>=end{
		return
	}
	tag:=arr[start]
	left:=start
	right:=end
	for left!=right{
		for arr[right]>=tag&&left<right{
			right--
		}
		for arr[left]<=tag&&left<right{
			left++
		}
		//划分左右半区，使比标记值小的都去左边，比它大的都去右边
		if left<right{
			arr[left],arr[right]=arr[right],arr[left]
		}
	}
	arr[start],arr[left]=arr[left],tag//换标记值至中部，左边都比它小，右边都比它大
	sort(arr,start,left-1)//左边
	sort(arr,right+1,end)//右边
}
func threeSumClosest(num []int,target int )int{
	sort(num,0,len(num)-1)
	abs:=func(n int)int{
		if n<0{
			return -n
		}else{
			return n
		}
	}

	res:=num[0]+num[1]+num[2]//最小的三数和
	tar:=abs(target-res)
	for i:=0;i<len(num)-2;i++ { //除第一个数以外，看剩下两个数可能的组合
		if i==2{
			c:=5
			c++
		}
		left := i + 1
		right := len(num) - 1
		for {
			if(right==left){
				break
			}
			resl := num[i] + num[left] + num[right]
			tarl := abs(target - resl)
			if tarl <= tar {
				tar = tarl
				res = resl
				//这是往中间靠，但并不是正确的移动方式
				//if(abs(num[left+1]-num[left])>abs(num[right]-num[right-1])){
				//	right--
				//}else{
				//	left++
				//}

			}
			if resl >target {
				right--
			} else if resl<target {
				left++
			}else{
				return target
			}
			}

		}



 	return res
}
func main(){
	n:=[]int{56,57,-47,-14,23,31,20,39,-51,7,-4,43,-53,32,24,56,-28,90,-75,-6,21,-100,41,-84,95,95,44,84,70,-22,-86,-6,90,-87,65,-28,-29,-94,98,-28,-100,23,-25,6,-56,-54,-5,53,-88,-25,-31,-71,-13,-62,73,-35,-78,16,99,97,84,-27,-43,-50,18,-16,-61,7,-17,16,-92,28,43,-38,-33,-27,84,-72,-100,-91,-97,-99,59,-63,73,99,98,-100,-37,-80,3,18,93,-81,12,-75,-43,99,10,10,-6,13,0,76,-82,-5,27,-38,-81,77,-55,-100,90,-32,-25,-15,-16,68,-6,87,65,-38,82,78,-61,87,-72,46,50,-60,86,39,69,85,-49,28}
	fmt.Println(threeSumClosest(n,-289))

}