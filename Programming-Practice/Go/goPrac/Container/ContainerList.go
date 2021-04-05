package main
import ("container/list"
	"fmt"
)

func print(l *list.List){
	for e:=l.Front();e!=nil;e=e.Next(){
		fmt.Println(e.Value)
	}
}
func main(){
	l:=list.New()
	l.PushBack(1)
	l.PushBack(2)
	print(l)
}
