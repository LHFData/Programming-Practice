package main

import (
	"os"
	"fmt"
	"net/http"
	"io"
)

type nodeInfo struct{
	//标示
	id string
	path string
	writer http.ResponseWriter
}
var nodeTable=make(map[string]string)

func main(){
	userId:=os.Args[1]
	fmt.Println(userId)

	nodeTable=map[string]string{
		"China":"localhost:1111",
		"Korea":"localhost:1112",
		"America":"localhost:1113",
		"Japan":"localhost:1114",
	}

	node:=nodeInfo{userId,nodeTable[userId],nil}
	fmt.Println("node"+node.id+" address:"+node.path+" is listenning")

	http.HandleFunc("/req",node.request)
	http.HandleFunc("/prePrepare",node.prePrepare)
	http.HandleFunc("/prepare",node.prepare)
	http.HandleFunc("/commit",node.commit)

	if err:=http.ListenAndServe(node.path,nil);err!=nil{
		fmt.Print(err)
	}

}
func (node *nodeInfo)request(writer http.ResponseWriter,request *http.Request){
	request.ParseForm()

	if(len(request.Form["warTime"])>0){
		node.writer=writer
		node.broadcast(request.Form["warTime"][0],"/prePrepare")
	}

}
func (node *nodeInfo)broadcast(msg string,path string){
	for nodeId,url:=range nodeTable{
		if nodeId==node.id{
			continue
		}
		http.Get("http://"+url+path+"?warTime="+msg+"&nodeId="+node.id)
	}
}
func (node *nodeInfo)prePrepare(writer http.ResponseWriter,request *http.Request){
	request.ParseForm()

	if(len(request.Form["warTime"])>0){
		node.broadcast(request.Form["warTime"][0],"/prepare")
	}
}
func(node *nodeInfo)prepare(writer http.ResponseWriter,request *http.Request){
	request.ParseForm()

	if len(request.Form["warTime"])>0{
		fmt.Println(request.Form["warTime"][0])
	}
	if len(request.Form["nodeId"])>0{
		fmt.Println(request.Form["nodeId"][0])
	}
	node.authentication(request)
}
var authenticationsuccess=true
var authenticationMap=make(map[string]string)

func(node *nodeInfo)authentication(request *http.Request){
	request.ParseForm()
	if authenticationsuccess!=false{
		if len(request.Form["nodeId"])>0{
			authenticationMap[request.Form["nodeId"][0]]="ok"
		}
	}
	if len(authenticationMap)>len(nodeTable)/3{
		node.broadcast(request.Form["warTime"][0],"/commit")
	}
}
func(node *nodeInfo)commit(writer http.ResponseWriter,request *http.Request){
	io.WriteString(node.writer,"达成共识")

}
