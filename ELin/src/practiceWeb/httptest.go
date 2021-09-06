package main

import (
	"fmt"
	"net/http"
)

//func main(){
//	//HandleFunc请求时调用对应的处理函数
//	http.HandleFunc("/",indexHandler)
//	http.HandleFunc("/hello",helloHandler)
//
//	log.Fatal(http.ListenAndServe(":9999",nil))
//}
func indexHandler(w http.ResponseWriter,req *http.Request){
	fmt.Fprintf(w,"URL.Path=%q\n",req.URL.Path)
}
func helloHandler(w http.ResponseWriter,req *http.Request){
	//for k,v :=range req.Header{
	//	//fmt.Fprintf(w,"Header[%q]=%q\n",k,v)
	//}
		fmt.Fprintf(w,"hello world")
}
