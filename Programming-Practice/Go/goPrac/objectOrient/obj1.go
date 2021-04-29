package main

import "fmt"

type shape interface {
	Area() int
}
type rectangle struct{
	length int
	width int
}
type circle struct{
	r int
}
func (r *rectangle)Area() int{
	return r.length*r.width
}
func (c *circle) Area() int{
	return c.r*c.r*3
}
func main(){
	r:=rectangle{2,3}
	c:=circle{3}
	shapes:=[]shape{&r,&c}
	for i:=0;i<2;i++{
		fmt.Println(shapes[i].Area())
	}
}
