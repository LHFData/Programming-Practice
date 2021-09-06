package main

import (
	"elin"
	"log"
	"net/http"
)

//请求处理的句柄
//type Engine struct{
//
//}

//func(engine *Engine)ServeHTTP(w http.ResponseWriter,req *http.Request){
//	switch req.URL.Path{
//	case "/":
//		fmt.Fprintf(w,"URL.Path=%q\n",req.URL.Path)
//	case "/hello":
//		for k,v:=range req.Header{
//		fmt.Fprintf(w,"header[%q]=%q",k,v)
//		}
//	default:
//		fmt.Fprintf(w,"404 NOT FOUND:%s\n",req.URL)
//
//	}
//
//}
func main(){
	engine:=new(elin.Engine)
	//统一的路由处理入口，可以针对性写路由处理逻辑和定义日志处理，异常处理
	log.Fatal(http.ListenAndServe(":9999",engine))
}