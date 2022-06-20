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

var methods = []string{methodGet, methodHead, methodPost, methodPut, methodPatch,
	methodDelete, methodConnect, methodOptions, methodTrace}

type Router struct {
	routers map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	rr := make(map[string]map[string]http.HandlerFunc)
	for _, m := range methods {
		rr[m] = make(map[string]http.HandlerFunc)
	}

	return &Router{
		routers: rr,
	}
}

func (r *Router) HandlerFunc(method, pattern string, handler http.HandlerFunc) {
	rr, ok := r.routers[method]
	if !ok {
		panic(fmt.Sprintf("register handler failed: the method(%s) error", method))
	}
	_, ok = rr[pattern]
	if ok {
		panic(fmt.Sprintf("the router path %s conflict", pattern))
	}
	r.routers[method][pattern] = handler
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

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path
	handler, ok := r.routers[method][path]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	handler(w, req)
}
