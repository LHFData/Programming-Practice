package main

import "fmt"

func romanToInt(s string) int {
	var output int
	for i:=0;i<len(s);i++{
		switch s[i] {
		case 'I':
			if i+1>=len(s){
				output+=1
				break
			}
			if s[i+1]=='V' {
				i++
				output+=4
			} else if s[i+1]=='X'{
				i++
				output+=9
			}else {
				output+=1
			}
		case 'X':
			if i+1>=len(s){
				output+=10
				break
			}
			if s[i+1]=='L'{
				i++
				output+=40
			} else if s[i+1]=='C'{
				i++
				output+=90
			}else {
				output+=10
			}
		case 'C':
			if i+1>=len(s){
				output+=100
				break
			}
			if s[i+1]=='D'{
				i++
				output+=400
			} else if s[i+1]=='M'{
				i++
				output+=900
			}else {
				output+=100
			}
		case 'V':output+=5
		case 'L':output+=50
		case 'D':output+=500
		case 'M':output+=1000
		}
	}
	return output
}

func main(){
	var input string
	fmt.Scan(&input)
	fmt.Println(romanToInt(input))
}