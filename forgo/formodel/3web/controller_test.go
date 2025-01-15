package formodel

import (
	"fmt"
	"testing"
)

/*
控制器模式（Controller Pattern）是一种用于将用户请求（input）转换为相应的响应（response）的软件设计模式。
	在 Web 应用程序中，控制器模式通常用于将用户请求转发给相应的处理程序，以及将处理程序的响应返回给用户。

	控制器模式可以帮助我们将请求与处理程序分离，从而使代码更加模块化和易于维护。
	在 Go 语言中，我们可以通过定义接口和结构体来实现控制器模式，并使用路由器将请求映射到相应的控制器。
*/

// Controller 定义一个控制器接口，包含一个处理请求的方法
type Controller interface {
	HandleRequest(request Request) Response
}

// Request 定义一个请求结构体，包含请求相关的信息
type Request struct {
	Path    string
	Method  string
	Payload []byte
}

// Response 定义一个响应结构体，包含响应相关的信息
type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       []byte
}

// Router 定义一个路由器结构体，用于将请求映射到响应的控制器
type Router struct {
	controllers map[string]Controller
}

// AddController 添加一个控制器到路由器中
func (r *Router) AddController(path string, controller Controller) {
	r.controllers[path] = controller
}

// HandleRequest 更具请求路径找到响应的控制器，并调用控制器的处理请求方法
func (r *Router) HandleRequest(request Request) Response {
	controller, ok := r.controllers[request.Path]
	if !ok {
		return Response{StatusCode: 404, Body: []byte("Not Found")}
	}
	return controller.HandleRequest(request)
}

// ExampleController 定义一个示例控制器，用于处理特定请求路径请求
type ExampleController struct{}

// HandleRequest 实现控制器接口中的处理请求方法
func (c *ExampleController) HandleRequest(request Request) Response {
	return Response{StatusCode: 200, Body: []byte("Hello, World!")}
}

func TestRouter(t *testing.T) {
	router := &Router{controllers: make(map[string]Controller)}
	router.AddController("/example", &ExampleController{})

	// 测试请求路径为 /example 的请求
	request := Request{Path: "/example", Method: "GET", Payload: nil}
	response := router.HandleRequest(request)
	fmt.Printf("response: code=%d, body=%s\n", response.StatusCode, string(response.Body))
}
