package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	ctx:=context.Background()
	ctx,cancelFunc:=context.WithTimeout(ctx,time.Second*4)
	defer cancelFunc()
	go targetFunc(ctx)
	for {
		select {
			case <-ctx.Done():
				switch ctx.Err() {
				case context.DeadlineExceeded:
					fmt.Println("context timeout exceeded")
					return
				case context.Canceled:
					fmt.Println("context cancelled by force")
					return
				}
		default:
			time.Sleep(time.Second*1)
			fmt.Println("SLEEP 1s")
		}
	}

}
func targetFunc(ctx context.Context){
	defer ctx.Done()
	time.Sleep(time.Second*3)
	fmt.Println("u r here")
}
