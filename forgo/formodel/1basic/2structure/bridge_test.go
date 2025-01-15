package formodel

import (
	"fmt"
	"testing"
)

/*
基本概念
桥接模式（Bridge Pattern）是一种结构型设计模式，它的UML图很像一座桥，它通过将【抽象部分】与【实现部分】分离，使它们可以独立变化，从而达到降低系统耦合度的目的。
桥接模式的主要目的是通过组合建立两个类之间的联系，而不是继承的方式。

举个简单的例子，图形编辑器中，每一种图形都需要蓝色、红色、黄色不同的颜色，
		如果不使用桥接模式，可能需要为每一种图形类型和每一种颜色都创建一个具体的子类，
		而使用桥接模式可以将图形和颜色两个维度分离，两个维度都可以独立进行变化和扩展，如果要新增其他颜色，只需添加新的 Color 子类，不影响图形类；反之亦然。

基本结构
桥接模式的基本结构分为以下几个角色：
	抽象Abstraction：一般是抽象类，定义抽象部分的接口，维护一个对【实现】的引用。
	修正抽象RefinedAbstaction：对抽象接口进行扩展，通常对抽象化的不同维度进行变化或定制。
	实现Implementor： 定义实现部分的接口，提供具体的实现。这个接口通常是抽象化接口的实现。
	具体实现ConcreteImplementor：实现实现化接口的具体类。这些类负责实现实现化接口定义的具体操作。

再举个例子，遥控器就是抽象接口，它具有开关电视的功能，修正抽象就是遥控器的实例，对遥控器的功能进行实现和扩展，
	而电视就是实现接口，具体品牌的电视机是具体实现，遥控器中包含一个对电视接口的引用，
	通过这种方式，遥控器和电视的实现被分离，我们可以创建多个遥控器，每个遥控器控制一个品牌的电视机，它们之间独立操作，不受电视品牌的影响，可以独立变化。

使用场景
桥接模式在日常开发中使用的并不是特别多，通常在以下情况下使用：
	当一个类存在两个独立变化的维度，而且这两个维度都需要进行扩展时，使用桥接模式可以使它们独立变化，减少耦合。
	不希望使用继承，或继承导致类爆炸性增长
总体而言，桥接模式适用于那些有多个独立变化维度、需要灵活扩展的系统。
*/

// 步骤1: 创建实现化接口
type TV interface {
	on()
	off()
	turnChannel()
}

// 步骤2: 创建具体实现化类
type ATV struct{}

func (atv *ATV) on() {
	fmt.Println("atv on")
}

func (atv *ATV) off() {
	fmt.Println("atv off")
}

func (atv *ATV) turnChannel() {
	fmt.Println("atv turnChannel")
}

type BTV struct{}

func (atv *BTV) on() {
	fmt.Println("btv on")
}

func (atv *BTV) off() {
	fmt.Println("btv off")
}

func (atv *BTV) turnChannel() {
	fmt.Println("btv turnChannel")
}

// 步骤3: 创建抽象化接口
type RemoteControl interface {
	PerformOperation()
}

// 步骤4: 创建扩充抽象化类
type PowerOperation struct {
	tv TV
}

func (po *PowerOperation) PerformOperation() {
	po.tv.on()
}

type OffOperation struct {
	tv TV
}

func (oo *OffOperation) PerformOperation() {
	oo.tv.off()
}

type ChannelSwitchOperation struct {
	tv TV
}

func (cso *ChannelSwitchOperation) PerformOperation() {
	cso.tv.turnChannel()
}

func Test_TVBridge(t *testing.T) {
	var tv TV
	tv = &ATV{}
	tv = &BTV{}

	var remoteControl RemoteControl
	remoteControl = &PowerOperation{tv: tv}
	remoteControl = &OffOperation{tv: tv}
	remoteControl = &ChannelSwitchOperation{tv: tv}
	remoteControl.PerformOperation()
}

/*
我们定义了一个 Sender 接口，以及两个实现了这个接口的结构体 SMS 和 Email。

接下来定义了一个抽象通知类 Notification，其中包含一个 Sender 接口类型的成员变量 sender。
Notification 类中定义了一个 SendNotification 方法，该方法将调用 sender 的 Send 方法，实现了消息发送的功能。

接着我们定义了两个具体的通知类 SMSNotification 和 EmailNotification。
这两个类都继承了 Notification 类，并在构造函数中分别传入了 SMS 和 Email 类型的 sender 对象，从而实现了不同的消息发送方式。

这样，我们就实现了一个简单的使用桥接模式的通知功能。通过抽象出消息发送功能，我们可以轻松地增加新的发送方式，而不需要修改原有的代码。
*/

type Sender interface {
	Send(message string)
}

type EmailSender struct{}
type SMSSender struct{}

func (es *EmailSender) Send(message string) {
	fmt.Println("EmailSender send message:", message)
}

func (ss *SMSSender) Send(message string) {
	fmt.Println("SMSSender send message:", message)
}

// 抽象通知类
type Notification struct {
	sender Sender
}

func (n *Notification) SendNotification(message string) {
	n.sender.Send(message)
}

// 修正抽象通知类
type SMSNotification struct {
	Notification
}

type EmailNotification struct {
	Notification
}

func NewSMSNotification(sender Sender) *SMSNotification {
	return &SMSNotification{Notification{sender: &SMSSender{}}}
}

func NewEmailNotification(sender Sender) *EmailNotification {
	return &EmailNotification{Notification{sender: &EmailSender{}}}
}

func Test_Bridge(t *testing.T) {
	smsSender := SMSSender{}
	emailSender := EmailSender{}

	smsNotification := NewSMSNotification(&smsSender)
	emailNotification := NewEmailNotification(&emailSender)

	smsNotification.SendNotification("hello world")
	emailNotification.SendNotification("hello world")
}
