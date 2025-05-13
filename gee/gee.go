package gee

import (
	"log"
	"net/http"
)

// HandlerFunc 定义一个新的类型 这个类型是函数 用于接收所有的HTTP请求并响应
type HandlerFunc func(c *Context)

// Engine 定义一个结构体 用于保存用户注册的路由和对应的处理函数
// 关于框架定义的一切HTTP操作都是通过Engine实例来调用的
type Engine struct {
	*RouterGroup // 匿名嵌套（继承RouterGroup的方法）
	router       *router
	groups       []*RouterGroup // 保存所有注册的路由分组
}

type RouterGroup struct {
	prefix      string        // 当前组的 URL 前缀，比如 "/api/v1"
	middlewares []HandlerFunc // 该组下要应用的中间件链
	parent      *RouterGroup
	engine      *Engine // 引擎指针，用于最终注册路由
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// 如果找到请求的方法和路径，则运行其定义的匿名函数，在这里就是handler
	//key := request.Method + "-" + request.URL.Path
	//if handler, ok := engine.router[key]; ok {
	//	handler(writer, request)
	//} else {
	//	_, _ = fmt.Fprintf(writer, "404 NOT FOUND: %s", request.URL)
	//}
	c := newContext(writer, request)
	// 原本的判断逻辑集成到router.go里了
	engine.router.handle(c)
}

func New() *Engine {
	// 构造函数 返回一个结构体地址作为框架的引擎
	// 相当于创建一个引擎实例
	engine := &Engine{
		router: newRouter(),
	}
	engine.RouterGroup = &RouterGroup{
		engine: engine,
	}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	// 给引擎实例新增路由、方法和对应操作
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(path string, handler HandlerFunc) {
	// 给引擎实例新增HTTP方法GET和对应操作
	group.addRoute("GET", path, handler)
}

func (group *RouterGroup) POST(path string, handler HandlerFunc) {
	// 给引擎实例新增HTTP方法POST和对应操作
	group.addRoute("POST", path, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	// 启动监听
	return http.ListenAndServe(addr, engine)
}
