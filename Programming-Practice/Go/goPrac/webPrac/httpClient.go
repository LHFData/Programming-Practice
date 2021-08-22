package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	resp,err:=http.Get("https://www.baidu.com")
	if err!=nil{
		fmt.Println("get failed,err:",err)
		return
	}
	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
	fmt.Println(body)
}
