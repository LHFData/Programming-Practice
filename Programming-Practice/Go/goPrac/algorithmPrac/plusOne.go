package main

import "fmt"

func plusOne(digits []int) []int {
	var add bool //进位
	current:=len(digits)-1
	_=add
	for current>=0{
		if digits[current]==9{
			digits[current]=0
			add=true
			current--
			continue
		}else{
			if add {
				if digits[current]+1==10{
					digits[current]=0
					continue
				}else{
					digits[current]=digits[current]+1
					add=false
					break
				}
			}else{
				digits[current]=digits[current]+1
				break
			}
		}
	}
	if add{
		digits=append(digits,0)
		digits[0]=1
	}
	return digits
}

func main()  {
	fmt.Println(plusOne([]int{8,9,9,9}))
}