package main

import "fmt"

type Minstack struct{
	min []int
	top int
	stack []int
}

func Constructor() Minstack{
	s:=new(Minstack)
	s.stack=append(s.stack,0)
	s.min=append(s.min,0)
	s.top=0
	return *s
}
func (this *Minstack)Push(x int){
	if this.top==0{
		this.stack[this.top]=x
		this.min[this.top]=x
		this.top++
		return
	}else if this.top>=len(this.stack){
		this.stack=append(this.stack,x)
		if x<this.min[this.top-1]{
			this.min=append(this.min,x)
		}else{
			this.min=append(this.min,this.min[this.top-1])
		}
		this.top++
		return
	}else{
		this.stack[this.top]=x
		if x<this.min[this.top-1]{
			this.min[this.top]=x
		}else{
			this.min[this.top]=this.min[this.top-1]
		}
		this.top++
		return
	}
}
func (this *Minstack) Pop(){
	if this.top==0{

	}else{
	this.top--
	}
}
func (this *Minstack)Top() int{
	return this.stack[this.top-1]
}
func (this *Minstack)Min() int{
	return this.min[this.top-1]

}
func main(){
	s:=Constructor()
	s.Push(1)
	s.Push(1)
	s.Push(2)


	s.Pop()
	s.Pop()
	s.Pop()
	s.Push(3)
	fmt.Println(s.Top())
	fmt.Println(s.Min())
	s.Push(-4)
	fmt.Println(s.Top())
	fmt.Println(s.Min())
	s.Pop()
	fmt.Println(s.Min())

}