
package main
//context是线程安全的，那么可以传进任意goroutine里面，假如有多个这样的goroutine，那么cancel就是一个总开关，可以一口气把多个goroutine停掉
import (
	"context"//上下文
	"log"
	"os"
	"time"
)

var logg *log.Logger
var ctx, cancel = context.WithCancel(context.Background())
func someHandler() {//可以作为103的起始

	go doStuff(ctx)
	go doStuff(ctx)

	//10秒后取消doStuff
	time.Sleep(10 * time.Second)


}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logg.Printf("done")
			return
		default:
			logg.Printf("work")
		}
	}
}

func main() {
	logg = log.New(os.Stdout, "", log.Ltime)
	someHandler()
	logg.Printf("111111")
	cancel()
	for ; ;  {
		time.Sleep(10*time.Second)
	}
}


