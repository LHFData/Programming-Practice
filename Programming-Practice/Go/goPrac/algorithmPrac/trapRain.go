package main

func trap(height []int) int {
	var left,right int
	var maxleft,maxright int
	res:=0
	left=0
	right=len(height)-1
	min:=func(a int,b int) int{
		if a<b{
			return a
		}
		return b
	}
	for left!=right{
		if height[left]>maxleft{
			maxleft=height[left]
		}
		if height[right]>maxright{
			maxright=height[right]
		}
		if height[left]<height[right]{
			res=res+min(maxleft,maxright)-height[left]
			left++
			continue
		}else if height[left]>=height[right]{
			res=res+min(maxleft,maxright)-height[right]
			right--
		}

	}
	return res
}
