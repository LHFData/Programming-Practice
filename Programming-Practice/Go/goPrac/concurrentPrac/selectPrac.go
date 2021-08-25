package main

import(
	"sync"
	"fmt"
	"time"
)

func publisher1(w *sync.WaitGroup,c chan<- int){
	defer w.Done()
	c<-1
	fmt.Println("pubulish 1")
}
func publisher2(w *sync.WaitGroup,c chan<- int){
	defer w.Done()
	c<-2
	fmt.Println("pubulish 2")
	time.Sleep(time.Second*2)
}
func reader(w *sync.WaitGroup,c <-chan int,cc <-chan int){
	defer w.Done()

L:
	for {
	select{
		case me:=<-c:
			fmt.Println(me)
		case he:=<-cc:
			fmt.Println(he)
			break L
		}
	}

}
func publisher3(w *sync.WaitGroup,cc chan<- int){
	defer w.Done()

	cc<-3

}
func main(){
	w:=&sync.WaitGroup{}
	c:=make(chan int,2)
	cc:=make(chan int)
	w.Add(4)
	go publisher1(w,c)
	go publisher2(w,c)
	go publisher3(w,cc)
	go reader(w,c,cc)
	w.Wait()
	defer close(c)
	defer func(){fmt.Println("close c")}()
	defer close(cc)
	defer func() {fmt.Println("close cc")}()
}