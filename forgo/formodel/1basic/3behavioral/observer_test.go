package formodel

import (
	"fmt"
	"testing"
)

/*
观察者模式
	观察者模式（发布-订阅模式）属于行为型模式，定义了一种一对多的依赖关系，让多个观察者对象同时监听一个主题对象，
	当主题对象的状态发生变化时，所有依赖于它的观察者都得到通知并被自动更新。

观察者模式依赖两个模块：
	Subject(主题)：也就是被观察的对象，它可以维护一组观察者，当主题本身发生改变时就会通知观察者。
	Observer(观察者)：观察主题的对象，当“被观察”的主题发生变化时，观察者就会得到通知并执行相应的处理。

使用观察者模式有很多好处，比如说
	1、观察者模式将主题和观察者之间的关系解耦，主题只需要关注自己的状态变化，而观察者只需要关注在主题状态变化时需要执行的操作，两者互不干扰。
	2、并且由于观察者和主题是相互独立的，可以轻松的增加和删除观察者，这样实现的系统更容易扩展和维护。

观察者模式依赖主题和观察者，但是一般有4个组成部分：
	主题Subject， 一般会定义成一个接口，提供方法用于注册、删除和通知观察者，通常也包含一个状态，当状态发生改变时，通知所有的观察者。
	观察者Observer: 观察者也需要实现一个接口，包含一个更新方法，在接收主题通知时执行对应的操作。
	具体主题ConcreteSubject: 主题的具体实现, 维护一个观察者列表，包含了观察者的注册、删除和通知方法。
	具体观察者ConcreteObserver: 观察者接口的具体实现，每个具体观察者都注册到具体主题中，当主题状态变化并通知到具体观察者，具体观察者进行处理。

什么时候使用观察者模式
	观察者模式特别适用于一个对象的状态变化会影响到其他对象，并且希望这些对象在状态变化时能够自动更新的情况。
	比如说在图形用户界面中，按钮、滑动条等组件的状态变化可能需要通知其他组件更新，这使得观察者模式被广泛应用于GUI框架，比如Java的Swing框架。
此外，观察者模式在前端开发和分布式系统中也有应用，比较典型的例子是前端框架Vue, 当数据发生变化时，视图会自动更新。
而在分布式系统中，观察者模式可以用于实现节点之间的消息通知机制，节点的状态变化将通知其他相关节点。
*/

type Observer interface {
	Update(subject ObSubject)
}

type ObSubject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	Notify()
}

type ConcreteObserver struct {
	name string
}

func (co *ConcreteObserver) Update(subject ObSubject) {
	fmt.Printf("%s received notification from %v\n", co.name, subject)
}

type ConcreteObSubject struct {
	observers []Observer
}

func (cs *ConcreteObSubject) Register(observer Observer) {
	cs.observers = append(cs.observers, observer)
}

func (cs *ConcreteObSubject) Unregister(observer Observer) {
	for i, o := range cs.observers {
		if o == observer {
			cs.observers = append(cs.observers[:i], cs.observers[i+1:]...)
			break
		}
	}
}

func (cs *ConcreteObSubject) Notify() {
	for _, o := range cs.observers {
		o.Update(cs)
	}
}

func TestObserver(t *testing.T) {
	// 创建主题和观察者
	subject := &ConcreteObSubject{}
	observer1 := &ConcreteObserver{name: "Observer 1"}
	observer2 := &ConcreteObserver{name: "Observer 2"}

	// 注册观察者
	subject.Register(observer1)
	subject.Register(observer2)

	// 发送通知
	subject.Notify()

	// 取消注册一个消费者
	subject.Unregister(observer1)

	// 发送另一个通知
	subject.Notify()
}
