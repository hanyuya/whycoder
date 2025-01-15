package formodel

import (
	"fmt"
	"testing"
)

/*
责任链模式是一种行为型设计模式，它允许你构建一个对象链，让请求从链的一端进入，然后沿着链上的对象依次处理，直到链上的某个对象能够处理该请求为止。
	职责链上的处理者就是一个对象，可以对请求进行处理或者将请求转发给下一个节点，
	这个场景在生活中很常见，就是一个逐层向上递交的过程，最终的请求要么被处理者所处理，要么处理不了，这也因此可能导致请求无法被处理。

责任链模式包括以下几个基本结构：
1. 处理者Handler ：定义一个处理请求的接口，包含一个处理请求的抽象方法和一个指向下一个处理者的链接。
2. 具体处理者ConcreteHandler: 实现处理请求的方法，并判断能否处理请求，如果能够处理请求则进行处理，否则将请求传递给下一个处理者。
3. 客户端：创建并组装处理者对象链，并将请求发送到链上的第一个处理者。

使用场景
责任链模式具有下面几个优点：
1. 降低耦合度：将请求的发送者和接收者解耦，每个具体处理者都只负责处理与自己相关的请求，客户端不需要知道具体是哪个处理者处理请求。
2. 增强灵活性：可以动态地添加或删除处理者，改变处理者之间的顺序以满足不同需求。

	但是由于一个请求可能会经过多个处理者，这可能会导致一些性能问题，并且如果整个链上也没有合适的处理者来处理请求，就会导致请求无法被处理。

	责任链模式是设计模式中简单且常见的设计模式，在日常中也会经常使用到，
		比如Java开发中过滤器的链式处理，以及Spring框架中的拦截器，都组装成一个处理链对请求、响应进行处理。
*/

type Handler interface {
	HandleRequest(request string)
	SetNext(handler Handler)
}

type BaseHandler struct {
	nextHandler Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.nextHandler = handler
}

type ConcreteHandler1 struct {
	BaseHandler
}

func (c *ConcreteHandler1) HandleRequest(request string) {
	if request == "A" {
		fmt.Println("ConcreteHandler1 can handle the request")
	} else if c.nextHandler != nil {
		c.nextHandler.HandleRequest(request)
	} else {
		fmt.Println("No handler cannot handle the request")
	}
}

type ConcreteHandler2 struct {
	BaseHandler
}

func (c *ConcreteHandler2) HandleRequest(request string) {
	if request == "B" {
		println("ConcreteHandler2 can handle the request")
	} else if c.nextHandler != nil {
		c.nextHandler.HandleRequest(request)
	} else {
		fmt.Println("No handler cannot handle the request")
	}
}

func TestChainOfResponsibility(t *testing.T) {
	ConcreteHandlerA := &ConcreteHandler1{}
	ConcreteHandlerB := &ConcreteHandler2{}
	ConcreteHandlerA.SetNext(ConcreteHandlerB)

	ConcreteHandlerA.HandleRequest("A")
	ConcreteHandlerA.HandleRequest("B")
	ConcreteHandlerA.HandleRequest("C")
}
