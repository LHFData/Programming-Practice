package elin

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

type HandlerFunc func(*Context)
type RouterGroup struct{
	prefix string
	middlewares []HandlerFunc
	parent *RouterGroup
	engine *Engine
}
//本质上分组是一棵树
type Engine struct {
	*RouterGroup //匿名字段，继承RouterGroup的所有方法
	router *router
	groups []*RouterGroup
	htmlTemplates *template.Template
	funcMap template.FuncMap
}
func (engine *Engine)SetFuncMap(funcMap template.FuncMap){
	engine.funcMap=funcMap
}
func (engine *Engine)LoadHTMLGlob(pattern string){
	engine.htmlTemplates=template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}
func New() *Engine{
	engine:=&Engine{router:newRouter()}
	engine.RouterGroup=&RouterGroup{engine: engine}
	return engine
}
func (group *RouterGroup) Group(prefix string) *RouterGroup{
	engine:=group.engine
	newGroup:=&RouterGroup{
		prefix: group.prefix+prefix,
		parent:group,
		engine: engine,
	}
	engine.groups=append(engine.groups,newGroup)
	return newGroup
}
func (group *RouterGroup)addRoute(method string,comp string,handler HandlerFunc)  {
	pattern:=group.prefix+comp
	log.Printf("Route %4s-%s",method,pattern)
	group.engine.router.addRoute(method,pattern,handler)
}


func (group *RouterGroup) GET(pattern string,handler HandlerFunc)  {
	group.addRoute("GET",pattern,handler)
}

func (group *RouterGroup) POST(pattern string,handler HandlerFunc){
	group.addRoute("POST",pattern,handler)
}
func (engine *Engine) RUN(addr string)(err error){
	return http.ListenAndServe(addr,engine)
}
func (group *RouterGroup)Use(middlewares ...HandlerFunc){
	group.middlewares=append(group.middlewares,middlewares...)
}
func (group *RouterGroup)createStaticHandler(relativePath string,fs http.FileSystem)HandlerFunc{
	absolutePath:=path.Join(group.prefix,relativePath)
	fileServer:=http.StripPrefix(absolutePath,http.FileServer(fs))
	return func(c *Context){
		file:=c.Param("filepath")
		if _,err:=fs.Open(file);err!=nil{
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer,c.Req)
	}
}
func(group *RouterGroup) Static(relativePath string,root string){
	handler:=group.createStaticHandler(relativePath,http.Dir(root))
	urlPattern:=path.Join(relativePath,"/filepath")
	group.GET(urlPattern,handler)
}
func (engine *Engine)ServeHTTP(w http.ResponseWriter,req *http.Request)  {
	var middlewares []HandlerFunc
	//通过前缀判断该请求适用哪些中间件
	//比如查询请求我们可以适用查询优化，缓存记录，修改请求可以适用修改日志记录等等
	//下面通过前缀判断，得到中间件列表以后，赋值给c。handlers
	for _,group:=range engine.groups{
		if strings.HasPrefix(req.URL.Path,group.prefix){
			middlewares=append(middlewares,group.middlewares...)
		}
	}
	c:=newContext(w,req)
	c.handlers=middlewares
	c.engine=engine
	engine.router.handle(c)
}