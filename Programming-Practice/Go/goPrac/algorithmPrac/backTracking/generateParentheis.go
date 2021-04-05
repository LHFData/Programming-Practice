package main

import "fmt"

var Res []string
func dfs(left int ,right int ,s string){
	if right==0&&left==0{
		Res = append(Res, s)
		//fmt.Println(s)
		return
	}
	if left>right{
		return
	}
	if left>0{
		s=s+"("
		dfs(left-1,right,s)
	}
	if right>0{
		s=s+")"
		dfs(left,right-1,s)
	}
}
func generateParenthesis(n int) []string {

	var s string

	dfs(n,n,s)

	return Res
}

func main(){
	fmt.Println(generateParenthesis(2))
}