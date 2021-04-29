package main

import "fmt"

func isValid(s string) bool {
	length:=len(s)
	stack:=make([]rune,length)
	top:=-1
	if length==0{
		return false
	}
	for i:=0;i<length;i++{
		if s[i]=='}'{
			if top==-1{
				return false
			}
			if stack[top]=='{'{
				stack[top]='#'
				top--
			}else{
				return false
			}
		}else if s[i]==']'{
			if top==-1{
				return false
			}
			if stack[top]=='['{
				stack[top]='#'
				top--
			}else{
				return false
			}
		}else if s[i]==')'{
			if top==-1{
				return false
			}
			if stack[top]=='('{
				stack[top]='#'
				top--
			}else{
				return false
			}
		}else{
			top++
			stack[top]= rune(s[i])

		}
	}
	if top>=0{
		return false
	}
	return true
}
func main (){


	fmt.Println(isValid(""))

}