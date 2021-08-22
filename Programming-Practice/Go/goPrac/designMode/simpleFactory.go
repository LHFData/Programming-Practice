package main
import "fmt"

type lhf interface {
	blackLhf()
}
type person struct{
	typeMe string
}
func(p person) blackLhf(){
	fmt.Println("blackLhf person")
}
type monster struct {
	typeMe string
}

func (m monster) blackLhf()  {
	fmt.Println("blackLhf monster")
}
func lhfShout(l lhf){
	l.blackLhf()
}
func main(){
	p:=person{typeMe: "person"}
	m:=monster{typeMe: "monster"}
	lhfShout(p)
	lhfShout(m)
	fmt.Println()
}