package formodel

import (
	"fmt"
	"testing"
)

/*
适配器模式Adapter是一种结构型设计模式，它可以将一个类的接口转换成客户希望的另一个接口，主要目的是充当两个不同接口之间的桥梁，使得原本接口不兼容的类能够一起工作。

应用场景
在开发过程中，适配器模式往往扮演者“补救”和“扩展”的角色：
	当使用一个已经存在的类，但是它的接口与你的代码不兼容时，可以使用适配器模式。
	在系统扩展阶段需要增加新的类时，并且类的接口和系统现有的类不一致时，可以使用适配器模式。
	使用适配器模式可以将客户端代码与具体的类解耦，客户端不需要知道被适配者的细节，客户端代码也不需要修改，这使得它具有良好的扩展性，但是这也势必导致系统变得更加复杂。

具体来说，适配器模式有着以下应用：
	不同的项目和库可能使用不同的日志框架，不同的日志框架提供的API也不同，因此引入了适配器模式使得不同的API适配为统一接口。
	Spring MVC中，HandlerAdapter 接口是适配器模式的一种应用。它负责将处理器（Handler）适配到框架中，使得不同类型的处理器能够统一处理请求。

	在.NET中，DataAdapter 用于在数据源（如数据库）和 DataSet 之间建立适配器，将数据从数据源适配到 DataSet 中，以便在.NET应用程序中使用。
*/

type Target interface {
	Request()
}

type Adaptee struct{}

func (a *Adaptee) SpecificRequest() {
	fmt.Println("Specific request")
}

type Adapter struct {
	adaptee *Adaptee
}

func (a *Adapter) Request() {
	a.adaptee.SpecificRequest()
}

func Test_Adapter(t *testing.T) {
	var adapter = &Adapter{
		adaptee: &Adaptee{},
	}
	adapter.Request()
}
