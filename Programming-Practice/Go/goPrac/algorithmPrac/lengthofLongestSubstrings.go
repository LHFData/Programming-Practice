package main

import "fmt"

func lengthofLongestSubstring(s string) int{
	var stringmap map[uint8]int
	stringmap=make(map[uint8]int)
	head:=0
	length:=0
	for i:=0;i<len(s);i++{
		_,ok :=stringmap[s[i]]
		if ok{
			head++
		}else{
			stringmap[s[i]]=i
			length++
		}
	}
	fmt.Println(length)
	fmt
}
