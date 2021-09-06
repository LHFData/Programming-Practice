package main

import (
	"elin"
	"net/http"
)

func main(){
	r:=elin.New()
	r.GET("/",func(c *elin.Context){
		c.HTML(http.StatusOK,"<h1>Hello Elin</h1>")
	})
	r.GET("/hello",func(c *elin.Context){
		c.String(http.StatusOK,"hello %s,u r at %s\n",c.Query("name"),c.Path)
	})
	r.GET("/hello:name", func(c *elin.Context) {
		c.String(http.StatusOK,"hello %s,u r at %s \n",c.Param("name"),c.Path)
	})
	r.GET("/assets/*filepath",func(c *elin.Context){
		c.JSON(http.StatusOK,elin.H{"filepath":c.Param("filepath")})
	})
	r.RUN(":9999")
}
