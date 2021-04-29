package main

import "fmt"

func main()  {
	m:=make(map[string]map[string]int)
	s:=make(map[string]int)
	s["l"]=1
	m["liaohaifeng"]=s
	res,ok:=m["liaohaifeng"]["l"]
	fmt.Println(ok,res)
	//g:=make(map[string]int)
	m["liaohfieng"]["h"]=make(map[string]int)


	res,ok=m["liaohaifeng"]["l"]
	fmt.Println(ok,res)
}
