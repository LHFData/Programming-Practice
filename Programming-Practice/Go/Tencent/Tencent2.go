package main

import (
	"fmt"
)

type stack struct{
	Length int
	Stack []int
	top int
	Cap int
}
func (s *stack)Init(length int){
	s.Length=length
	s.Stack=make([]int,s.Length)
	s.top=0
	s.Cap=0
}
func (s *stack)Push(num int){
	if s.top>0 {
		if num+s.Stack[s.top-1] == 10 {
			s.top--
			s.Cap--
		} else {
			s.Stack[s.top] = num
			s.top++
			s.Cap++
		}
	} else{
		s.Stack[s.top] = num
		s.top++
		s.Cap++
	}
}
func main(){
	var s stack
	s.Init(10)
	s.Push(2)
	s.Push(1)
	s.Push(3)
	s.Push(7)
	s.Push(9)
	s.Push(2)
	fmt.Println(s.Cap)
}