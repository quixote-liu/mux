package mux

import (
	"fmt"
	"net/http"
)

var methods = []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch,
	http.MethodDelete, http.MethodConnect, http.MethodOptions, http.MethodTrace}

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
	r.HandlerFunc(http.MethodGet, pattern, handler)
}

func (r *Router) POST(pattern string, handler http.HandlerFunc) {
	r.HandlerFunc(http.MethodPost, pattern, handler)
}

func (r *Router) PUT(pattern string, handler http.HandlerFunc) {
	r.HandlerFunc(http.MethodPut, pattern, handler)
}

func (r *Router) PATCH(pattern string, handler http.HandlerFunc) {
	r.HandlerFunc(http.MethodPatch, pattern, handler)
}

func (r *Router) DELETE(pattern string, handler http.HandlerFunc) {
	r.HandlerFunc(http.MethodDelete, pattern, handler)
}

func (r *Router) OPTIONS(pattern string, handler http.HandlerFunc) {
	r.HandlerFunc(http.MethodOptions, pattern, handler)
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
