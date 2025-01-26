package formodel

import (
	"fmt"
	"testing"
)

/*
享元模式是一种结构型设计模式，在享元模式中，对象被设计为可共享的，可以被多个上下文使用，而不必在每个上下文中都创建新的对象。
想要了解享元模式，就必须要区分什么是内部状态，什么是外部状态。
	内部状态是指那些可以被多个对象共享的状态，它存储在享元对象内部，并且对于所有享元对象都是相同的，这部分状态通常是不变的。
	外部状态是享元对象依赖的、可能变化的部分。这部分状态不存储在享元对象内部，而是在使用享元对象时通过参数传递给对象。

例子:假设我们在构建一个简单的图形编辑器，用户可以在画布上绘制不同类型的图形，
	而图形就是所有图形对象的内部状态（不变的），而图形的坐标位置就是图形对象的外部状态（变化的）。

如果图形编辑器中有成千上万的图形对象，每个图形对象都独立创建并存储其内部状态，那么系统的内存占用可能会很大，
	在这种情况下，享元模式共享相同类型的图形对象，每种类型的图形对象只需创建一个共享实例，然后通过设置不同的坐标位置个性化每个对象，
	通过共享相同的内部状态，降低了对象的创建和内存占用成本。

享元模式包括以下几个重要角色：
	享元接口Flyweight: 所有具体享元类的共享接口，通常包含对外部状态的操作。
	具体享元类ConcreteFlyweight: 继承Flyweight类或实现享元接口，包括内部状态和外部状态。内部状态在对象创建时就被固定，而外部状态则在对象使用时被传入。
	享元工厂类FlyweightFactory: 创建并管理享元对象，当用户请求时，提供已创建的实例或者创建一个。
	客户端Client: 维护外部状态，在使用享元对象时，将外部状态传递给享元对象。

使用场景
	使用享元模式的关键在于包含大量相似对象，并且这些对象的内部状态可以共享。
	具体的应用场景包括文本编辑器，图形编辑器，游戏中的角色创建，这些对象的内部状态比较固定(外观，技能，形状)，但是外部状态变化比较大时，可以使用。
*/

type Flyweight interface {
	Operation(extrinsicState string)
}

type ConcreteFlyweight struct {
	intrinsicState string
}

func (f *ConcreteFlyweight) Operation(extrinsicState string) {
	fmt.Printf("ConcreteFlyweight: intrinsic state is %s, extrinsic state is %s\n", f.intrinsicState, extrinsicState)
}

type FlyweightFactory struct {
	flyweights map[string]Flyweight
}

func (factory *FlyweightFactory) GetFlyweight(key string) Flyweight {
	if flyweight, ok := factory.flyweights[key]; ok {
		return flyweight
	}
	flyweight := &ConcreteFlyweight{intrinsicState: key}
	factory.flyweights[key] = flyweight
	return flyweight
}

func TestFlyweight(t *testing.T) {
	factory := &FlyweightFactory{flyweights: make(map[string]Flyweight)}
	flyweight1 := factory.GetFlyweight("flyweight1")
	flyweight2 := factory.GetFlyweight("flyweight2")
	flyweight3 := factory.GetFlyweight("flyweight3")
	flyweight4 := factory.GetFlyweight("flyweight4")
	flyweight5 := factory.GetFlyweight("flyweight5")

	flyweight1.Operation("external state 1")
	flyweight2.Operation("external state 2")
	flyweight3.Operation("external state 3")
	flyweight4.Operation("external state 4")
	flyweight5.Operation("external state 5")
}
