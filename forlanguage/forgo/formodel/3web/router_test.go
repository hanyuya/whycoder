package formodel

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

/*
路由器模式（Router Pattern）指的是一种将 URL 路径与相应的处理函数进行映射的设计模式。
	通过路由器模式，我们可以实现基于 URL 路径的请求分发，将不同的请求路由到不同的处理函数中进行处理。

在实际应用中，我们可能需要实现更复杂的路由匹配规则，支持参数化路径、中间件等功能。
	但是，了解基本的路由器模式原理之后，我们就可以很方便地实现自己的路由器，从而实现基于 URL 路径的请求分发。
*/

// 定义路由器结构体
type MyRouter struct {
	handlers map[string]http.HandlerFunc
}

// 定义路由器的构造函数
func NewMyRouter() *MyRouter {
	return &MyRouter{
		handlers: make(map[string]http.HandlerFunc),
	}
}

// HandleFunc 向路由器中添加新的处理函数
func (r *MyRouter) AddHandler(path string, handler http.HandlerFunc) {
	r.handlers[path] = handler
}

// ServeHTTP 实现路由器的路由功能
func (r *MyRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 获取请求的 URL 路径
	path := req.URL.Path

	// 从路由器中查找对应的处理函数
	handler, ok := r.handlers[path]
	if !ok {
		// 如果没有找到对应的处理函数，则返回 404 错误
		http.NotFound(w, req)
		return
	}

	// 调用处理函数处理请求
	handler(w, req)
}

// 处理 / 路径的请求
func IndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Index Page"))
}

// 处理 /about 路径的请求
func AboutHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("About Page"))
}

// 处理 /contact 路径的请求
func ContactHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Contact Page"))
}

func Test_MyRouter(t *testing.T) {
	// 创建路由器
	router := NewMyRouter()

	// 向路由器中添加处理函数
	router.AddHandler("/", IndexHandler)
	router.AddHandler("/about", AboutHandler)
	router.AddHandler("/contact", ContactHandler)

	req := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Path: "/",
		},
	}
	resp := httptest.NewRecorder()

	// 调用 ServeHTTP 方法处理请求
	router.ServeHTTP(resp, req)

	fmt.Printf("%s\n", resp.Body.String())

}
