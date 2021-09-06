package main
//
//import (
//	"fmt"
//	"net/http"
//)
//
//type HandlerFunc func(http.ResponseWriter,*http.Request)
//
//type Engine struct{
//	router map[string]HandlerFunc
//}
//func New()*Engine{
//	return &Engine{router:make(map[string]HandlerFunc)}
//}
//func (engine *Engine) addRoute(method string,pattern string,handler HandlerFunc){
//	key:=method+"-"+pattern
//	engine.router[key]=handler
//	//根据请求方法调用对应的处理程序
//}
//
//func(engine *Engine) GET(pattern string,handler HandlerFunc){
//	engine.addRoute("GET",pattern,handler)
//}
//func (engine *Engine)POST(pattern string,handler HandlerFunc)  {
//	engine.addRoute("POST",pattern,handler)
//}
//func (engine *Engine)RUN(addr string)(err error){
//	return http.ListenAndServe(addr,engine)
//}
////实现Handler接口的ServeHTTP方法
//func (engine *Engine)ServeHTTP(w http.ResponseWriter,req *http.Request){
//	key:=req.Method+"-"+req.URL.Path
//	if handler,ok:=engine.router[key];ok{
//		handler(w,req)
//	}else{
//		fmt.Fprintf(w,"404 NOT FOUND %q",req.URL)
//	}
//}


