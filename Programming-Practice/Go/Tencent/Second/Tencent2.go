package main

import "fmt"

func same(s1 string,s2 string)bool{
	for i:=0;i<len(s1);i++{
		if s1[i]!=s2[i]{
			return false
		}
	}
	return true
}
func sameReverse(s1 string,s2 string) bool{
	if len(s1)==0{
		return true
	}
	if s1==s2{
		return true
	}else{
		l:=len(s1)
		if sameReverse(s1[:l/2],s2[:l/2])||sameReverse(s1[:l/2],s2[l/2:])||sameReverse(s1[l/2:],s2[:l/2])||sameReverse(s1[l/2:],s2[l/2:]){
			return true
		}else{
			return false
		}
	}
}
func main(){

	var n int
	fmt.Scan(&n)
	var compareA []string
	var compareB []string
	for i:=0;i<n;i++{
		var temp1 string
		var temp2 string
		fmt.Scan(&temp1)
		fmt.Scan(&temp2)
		compareA=append(compareA,temp1)
		compareB=append(compareB,temp2)
	}
	for i:=0;i<n;i++{
		if len(compareA[i])%2==0{
			l:=len(compareA[i])

			if sameReverse(compareA[i][:l/2],compareB[i][:l/2])||sameReverse(compareA[i][:l/2],compareB[i][l/2:])||sameReverse(compareA[i][l/2:],compareB[i][:l/2])||sameReverse(compareA[i][l/2:],compareB[i][l/2:]){
				fmt.Println("YES")
			}else{
				fmt.Println("NO")
			}
		}else{
			if same(compareA[i],compareB[i]){
				fmt.Println("YES")
			}else{
				fmt.Println("NO")
			}

		}
	}
}
