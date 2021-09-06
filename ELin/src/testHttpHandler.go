package main
import (
	"fmt"
	"net/http"
	"elin"
)
func main(){
	r:= elin.New()
	r.GET("/hello",func(w http.ResponseWriter,req *http.Request){
		fmt.Fprintf(w,"URL.Path=%q\n",req.URL.Path)
	})
	r.RUN(":9999")
}


