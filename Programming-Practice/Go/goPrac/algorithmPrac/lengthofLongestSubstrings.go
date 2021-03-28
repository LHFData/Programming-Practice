package main

import "fmt"

func lengthofLongestSubstring(s string) int{
	var stringmap map[uint8]int
	stringmap=make(map[uint8]int)
	head:=0
	tail:=0
	length:=0
	for i:=0;i<len(s);i++{
		_,ok :=stringmap[s[i]]
		if ok{
			length=stringmap[s[i]]+1-head
			head=stringmap[s[i]]+1
			stringmap[s[i]]=i
		}else{
			stringmap[s[i]]=i
			tail++
			length++
		}
	}
	fmt.Println(length)
	fmt.Println(tail-head)
	return length

}
func main(){
	s:="accasdads"
	lengthofLongestSubstring(s)
}
