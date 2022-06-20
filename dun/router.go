package dun

import (
	"net/http"
	"strings"
)

type Router struct {
	roots    map[string]*Node
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{
		roots:    make(map[string]*Node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parse(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (router *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parse(pattern)
	key := method + "-" + pattern
	_, ok := router.roots[method]
	if !ok {
		router.roots[method] = &Node{}
	}
	router.roots[method].insert(pattern, parts, 0)
	router.handlers[key] = handler
}

func (router *Router) getRoute(method string, path string) (*Node, map[string]string) {
	searchParts := parse(path)
	params := make(map[string]string)
	root, ok := router.roots[method]
	if !ok {
		return nil, nil
	}
	node := root.search(searchParts, 0)
	if node == nil {
		return nil, nil
	}
	parts := parse(node.pattern)
	for index, part := range parts {
		if part[0] == ':' {
			params[part[1:]] = searchParts[index]
		}
		if part[0] == '*' && len(part) > 1 {
			params[part[1:]] = strings.Join(searchParts[index:], "/")
			break
		}
	}
	return node, params
}

func (router *Router) getRoutes(method string) []*Node {
	root, ok := router.roots[method]
	if !ok {
		return nil
	}
	nodes := make([]*Node, 0)
	root.travel(&nodes)
	return nodes
}

func (router *Router) handle(context *Context) {
	node, params := router.getRoute(context.Method, context.Path)
	if node != nil {
		context.Params = params
		key := context.Method + "-" + node.pattern
		router.handlers[key](context)
	} else {
		context.String(http.StatusNotFound, "404 NOT FOUND: %s\n", context.Path)
	}
}
