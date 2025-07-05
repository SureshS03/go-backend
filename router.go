package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"github.com/SureshS03/goconnect/internal/redis"
)

type router struct {
	routes map[string]map[string]http.HandlerFunc
}

func NewRouter(db *sql.DB) *router {
	r := &router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}
	//ADD wrap chain to all and test
	s := NewService(db)
	r.addRoute("POST", "/user/", WrapChain(s.AddUser, LogRequestMiddleware, SecureHeadersMiddleware))
	r.addRoute("GET", "/user/", s.GetAllUser)
	r.addRoute("GET", "/user/:id/", s.GetUser)
	r.addRoute("POST", "/posts/", s.Addpost)
	r.addRoute("GET", "/posts/user/:user_id", s.GetUserPost)
	r.addRoute("GET", "/posts/:id", s.GetPost)
	r.addRoute("GET", "/tokengen/", s.GenToken)

	return r
}

func (r *router) addRoute(method, path string, handler http.HandlerFunc) {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]http.HandlerFunc)
	}
	r.routes[method][path] = handler
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	methodRoutes := r.routes[req.Method]
	if methodRoutes == nil {
		http.NotFound(w, req)
		fmt.Println("Unknown HTTP method:", req.Method)
		return
	}

	if handler, ok := methodRoutes[req.URL.Path]; ok {
		handler(w, req)
		return
	}

	for routePath, handler := range methodRoutes {
		if match, params := matchPath(routePath, req.URL.Path); match {
			req = addParamsToRequest(req, params)
			handler(w, req)
			return
		}
	}

	http.NotFound(w, req)
	fmt.Println("No matching route for", req.URL.Path)
}

func matchPath(routePath, reqPath string) (bool, map[string]string) {
	routeParts := strings.Split(strings.Trim(routePath, "/"), "/")
	reqParts := strings.Split(strings.Trim(reqPath, "/"), "/")

	if len(routeParts) != len(reqParts) {
		return false, nil
	}

	params := make(map[string]string)
	for i := range reqParts {
		fmt.Println(i)
		if strings.HasPrefix(routeParts[i], ":") {
			key := strings.TrimPrefix(routeParts[i], ":")
			params[key] = reqParts[i]
			fmt.Println(params)
		} else if routeParts[i] != reqParts[i] {
			return false, nil
		}
	}
	return true, params
}

type contextKey string

func addParamsToRequest(req *http.Request, params map[string]string) *http.Request {
	ctx := req.Context()
	for k, v := range params {
		ctx = context.WithValue(ctx, contextKey(k), v)
	}
	return req.WithContext(ctx)
}

func GetParam(r *http.Request, key string) string {
	val := r.Context().Value(contextKey(key))
	if str, ok := val.(string); ok {
		return str
	}
	return ""
}

func WrapChain(h http.HandlerFunc, mws ...func(http.Handler) http.Handler) http.HandlerFunc {
	final := http.Handler(http.HandlerFunc(h))
	for _, mw := range mws {
		final = mw(final)
	}
	return final.ServeHTTP
}


func LogRequestMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("LOG %s - %s %s %s\n", r.RemoteAddr, r.Proto, r.Method, r.URL)
    
    next.ServeHTTP(w, r)
  })
}

func SecureHeadersMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("X-XSS-Protection", "1; mode-block")
    w.Header().Set("X-Frame-Options", "deny")
    
    next.ServeHTTP(w, r)
  })
}

func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		bearerToken := req.Header.Get("Authorization")
		token := strings.Split(bearerToken, "")[1]
		fmt.Println("token in req", token)
		rdToken, err := redis.GetCache("token:" + token)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "UnAuthorization", 600)
		}
		if token != rdToken {
			http.Error(w, "UnAuthorization", 600)
		}
		next.ServeHTTP(w, req)
	})
}
