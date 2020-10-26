package main

import (
	"fmt"
)

func main()  {
	defer_call()
}
//defer的使用遵循先进后出，即栈序，定义早执行慢。
//同时以定义时的变量值为准，而不管定义后数值的变化。
func defer_call()  {
	defer func(){fmt.Println("打印前")}()
	defer func() {fmt.Println("打印中")}()
	defer func() {fmt.Println("打印后")}()
	//panic("触发异常")

}