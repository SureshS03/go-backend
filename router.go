package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

type router struct{
	routes map[string]map[string]http.HandlerFunc
}

func NewRouter(db *sql.DB) *router {
	r:= &router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}

	s := NewService(db)
	r.addNewRoutes("POST", "/user", s.AddUser)
	r.addNewRoutes("GET", "/user", s.GetAllUser)
	r.addNewRoutes("GET", "/user/:id", s.GetUser)
	return r
}

func (r *router) addNewRoutes(method string, path string, handlerFunction http.HandlerFunc)  {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]http.HandlerFunc)
	}
	r.routes[method][path] = handlerFunction
	fmt.Println(r.routes)
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := r.routes[req.Method]
	if method == nil {
		http.NotFound(w, req)
		fmt.Println("non used method is called")
		return
	}

	if handler, ok := method[req.URL.Path]; ok {
		handler(w, req)
		return
	}

	for routePath, handlers := range method{
		fmt.Println(routePath)
		if match, params := matchPath(routePath, req.URL.Path); match{

		}
	}
}

func matchPath(routePath, reqPath string) (bool, error) {
	routePart := strings.Split(strings.Trim(routePath, "/"), "/")
	reqPart := strings.Split(strings.Trim(reqPath, "/"), "/")

	if len(routeParts) != len(actualParts) {
		return false, nil
	}

	for i:= range reqPart{

	}
}