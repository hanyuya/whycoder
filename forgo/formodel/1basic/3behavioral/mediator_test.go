package formodel

import (
	"fmt"
	"testing"
)

/*
中介者模式（Mediator Pattern）
	中介者模式（Mediator Pattern）也被称为调停者模式，是一种行为型设计模式，
		它通过一个中介对象来封装一组对象之间的交互，从而使这些对象不需要直接相互引用。这样可以降低对象之间的耦合度，使系统更容易维护和扩展。

	当一个系统中的对象有很多且多个对象之间有复杂的相互依赖关系时，依赖关系很难理清，这时我们可以引入一个中介者对象来进行协调和交互。
		中介者模式可以使得系统的网状结构变成以中介者为中心的星形结构，每个具体对象不再通过直接的联系与另一个对象发生相互作用，而是通过“中介者”对象与另一个对象发生相互作用。
		当一个对象发生改变时，它并不直接通知其他对象，而是通过中介者对象来通知，由中介者对象来负责将这个变化传递给其他对象。

基本结构
	中介者模式包括以下几个重要角色：
		抽象中介者（Mediator）： 定义中介者的接口，用于各个具体同事对象之间的通信。
		具体中介者（Concrete Mediator）： 实现抽象中介者接口，负责协调各个具体同事对象的交互关系，它需要知道所有具体同事类，并从具体同事接收消息，向具体同事对象发出命令。
		抽象同事类（Colleague）： 定义同事类的接口，维护一个对中介者对象的引用，用于通信。
		具体同事类（Concrete Colleague）： 实现抽象同事类接口，每个具体同事类只知道自己的行为，而不了解其他同事类的情况，因为它们都需要与中介者通信，通过中介者协调与其他同事对象的交互。

使用场景
	中介者模式使得同事对象不需要知道彼此的细节，只需要与中介者进行通信，简化了系统的复杂度，也降低了各对象之间的耦合度，
		但是这也会使得中介者对象变得过于庞大和复杂，如果中介者对象出现问题，整个系统可能会受到影响。

	中介者模式适用于当系统对象之间存在复杂的交互关系或者系统需要在不同对象之间进行灵活的通信时使用，可以使得问题简化，
*/

// 中介者接口
type Mediator interface {
	sendMessage(msg string, user User)
	receiveMessage() string
}

// 具体中介者
type ConcreteMediator struct {
	Message string
}

func (cm *ConcreteMediator) sendMessage(msg string, user User) {
	cm.Message = fmt.Sprintf("%s 发送消息：%s\n", user.getName(), msg)
}

func (cm *ConcreteMediator) receiveMessage() string {
	return cm.Message
}

type User struct {
	name     string
	mediator Mediator
}

func (u *User) getName() string {
	return u.name
}

func (u *User) setMediator(mediator Mediator) {
	u.mediator = mediator
}

func (u *User) sendMessage(msg string) {
	u.mediator.sendMessage(msg, *u)
}

func (u *User) receiveMessage() string {
	return u.mediator.receiveMessage()
}

func TestMediator(t *testing.T) {
	mediator := ConcreteMediator{}
	user1 := User{"user1", &mediator}
	user2 := User{"user2", &mediator}
	user1.setMediator(&mediator)
	user2.setMediator(&mediator)
	user1.sendMessage("hello")
	fmt.Println(user2.receiveMessage())
}
