package mux

import (
	"fmt"
	"net/http"
)

const (
	methodGet     = http.MethodGet
	methodHead    = http.MethodHead
	methodPost    = http.MethodPost
	methodPut     = http.MethodPut
	methodPatch   = http.MethodPatch
	methodDelete  = http.MethodDelete
	methodConnect = http.MethodConnect
	methodOptions = http.MethodOptions
	methodTrace   = http.MethodTrace
)

var (
	methods = []string{methodGet, methodHead, methodPost, methodPut, methodPatch,
		methodDelete, methodConnect, methodOptions, methodTrace}
)

type Router struct {
	groups map[string]*Group
}

func NewRouter() *Router {
	groups := make(map[string]*Group)
	for _, m := range methods {
		groups[m] = newGroup()
	}

	return &Router{
		groups: groups,
	}
}

// func (r *Router) Group(prefix string) *Group {

// }

func (r *Router) HandlerFunc(method, pattern string, handler http.HandlerFunc) {
	g, ok := r.groups[method]
	if !ok {
		panic(fmt.Sprintf("register handler failed: the method(%s) error", method))
	}
	if err := g.register(pattern, handler); err != nil {
		panic(fmt.Sprintf("register handler failed: %v", err))
	}
}

func (r *Router) GET(pattern string, handler http.HandlerFunc) {
	r.HandlerFunc(methodGet, pattern, handler)
}

func (r *Router) POST(pattern string, handler http.HandlerFunc) {
	r.HandlerFunc(methodPost, pattern, handler)
}

func (r *Router) PUT(pattern string, handler http.HandlerFunc) {
	r.HandlerFunc(methodPut, pattern, handler)
}

func (r *Router) PATCH(pattern string, handler http.HandlerFunc) {
	r.HandlerFunc(methodPatch, pattern, handler)
}

func (r *Router) DELETE(pattern string, handler http.HandlerFunc) {
	r.HandlerFunc(methodDelete, pattern, handler)
}

func (r *Router) OPTIONS(pattern string, handler http.HandlerFunc) {
	r.HandlerFunc(methodOptions, pattern, handler)
}

func (r *Router) Any(pattern string, handler http.HandlerFunc) {
	for _, m := range methods {
		r.HandlerFunc(m, pattern, handler)
	}
}
