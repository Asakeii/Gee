package gee

import (
	"log"
	"net/http"
)

type router struct {
	routes map[string]HandlerFunc
}

func newRouter() *router {
	return &router{make(map[string]HandlerFunc)}
}

func (router *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route  %4s - %s\n", method, pattern)
	router.routes[method+"-"+pattern] = handler
}

func (router *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := router.routes[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 page not found")
	}
}
