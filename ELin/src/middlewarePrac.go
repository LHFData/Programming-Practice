package main

import (
	"elin"
	"log"
	"net/http"
	"time"
)

func onlyForV2()elin.HandlerFunc{
	return func(c *elin.Context){
	t:=time.Now()
	c.Fail(500,"Internal Server error")
	log.Printf("[%d] %s in %v for group v2",c.StatusCode,c.Req.RequestURI,time.Since(t))
	}
}
func main(){
	r:=elin.New()
	r.Use(elin.Logger())
	r.GET("/",func(c *elin.Context){
		c.HTML(http.StatusOK,"<h1>hello elin<h2>")
	})
	v2:=r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *elin.Context) {
			c.String(http.StatusOK,"hello %s,you re at %s\n",c.Param("name"),c.Path)
		})
	}
	r.RUN(":9999")
}
