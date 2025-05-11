package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc 定义一个新的类型 这个类型是函数 用于接收所有的HTTP请求并响应
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 定义一个结构体 用于保存用户注册的路由和对应的处理函数
// 关于框架定义的一切HTTP操作都是通过Engine实例来调用的
type Engine struct {
	router map[string]HandlerFunc
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// 如果找到请求的方法和路径，则运行其定义的匿名函数，在这里就是handler
	key := request.Method + "-" + request.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(writer, request)
	} else {
		_, _ = fmt.Fprintf(writer, "404 NOT FOUND: %s", request.URL)
	}
}

func New() *Engine {
	// 构造函数 返回一个结构体地址作为框架的引擎
	// 相当于创建一个引擎实例
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

func (engine *Engine) addRoute(method string, path string, handler HandlerFunc) {
	// 给引擎实例新增路由、方法和对应操作
	engine.router[method+"-"+path] = handler
}

func (engine *Engine) GET(path string, handler HandlerFunc) {
	// 给引擎实例新增HTTP方法GET和对应操作
	engine.addRoute("GET", path, handler)
}

func (engine *Engine) POST(path string, handler HandlerFunc) {
	// 给引擎实例新增HTTP方法POST和对应操作
	engine.addRoute("POST", path, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	// 启动监听
	return http.ListenAndServe(addr, engine)
}
