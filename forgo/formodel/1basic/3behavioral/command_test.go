package formodel

import "testing"

/*
命令模式

	命令模式是一种行为型设计模式，其允许将请求封装成一个对象(命令对象，包含执行操作所需的所有信息)，
		并将命令对象按照一定的顺序存储在队列中，然后再逐一调用执行，这些命令也可以支持反向操作，进行撤销和重做。

	这样一来，发送者只需要触发命令就可以完成操作，不需要知道接受者的具体操作，从而实现两者间的解耦。

	举个现实中的应用场景，遥控器可以控制不同的设备，
		在命令模式中，可以假定每个按钮都是一个命令对象，包含执行特定操作的命令，不同设备对同一命令的具体操作也不同，这样就可以方便的添加设备和命令对象。

基本结构
	命令模式包含以下几个基本角色：
		命令接口Command：接口或者抽象类，定义执行操作的接口。
		具体命令类ConcreteCommand: 实现命令接口，执行具体操作，在调用execute方法时使“接收者对象”根据命令完成具体的任务，比如遥控器中的“开机”，“关机”命令。
		接收者类Receiver: 接受并执行命令的对象，可以是任何对象，遥控器可以控制空调，也可以控制电视机，电视机和空调负责执行具体操作，是接收者。
		调用者类Invoker: 发起请求的对象，有一个将命令作为参数传递的方法。它不关心命令的具体实现，只负责调用命令对象的 execute() 方法来传递请求，在本例中，控制遥控器的“人”就是调用者。
		客户端：创建具体的命令对象和接收者对象，然后将它们组装起来。

优缺点和使用场景
	命令模式在需要将请求封装成对象、支持撤销和重做、设计命令队列等情况下，都是一个有效的设计模式。
		撤销操作： 需要支持撤销操作，命令模式可以存储历史命令，轻松实现撤销功能。
		队列请求： 命令模式可以将请求排队，形成一个命令队列，依次执行命令。
		可扩展性： 可以很容易地添加新的命令类和接收者类，而不影响现有的代码。新增命令不需要修改现有代码，符合开闭原则。
	但是对于每个命令，都会有一个具体命令类，这可能导致类的数量急剧增加，增加了系统的复杂性。

	命令模式同样有着很多现实场景的应用，比如Git中的很多操作，如提交（commit）、合并（merge）等，都可以看作是命令模式的应用，用户通过执行相应的命令来操作版本库。
	Java的GUI编程中，很多事件处理机制也都使用了命令模式。例如，每个按钮都有一个关联的 Action，它代表一个命令，按钮的点击触发 Action 的执行。
*/

// 命令接口Command
type Command interface {
	Execute()
}

// 具体命令类ConcreteCommand: 实现命令接口，执行具体操作，
// 在调用execute方法时使“接收者对象”根据命令完成具体的任务，比如遥控器中的“开机”，“关机”命令。
type ConcreteCommand struct {
	receiver Receiver
}

func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}

type Receiver interface {
	Action()
}

type Receiver1 struct {
	name string
}

func (r *Receiver1) Action() {
	println("Receiver1 action")
}

type Receiver2 struct {
	name string
}

func (r *Receiver2) Action() {
	println("Receiver2 action")
}

type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func TestCommand(t *testing.T) {
	receiver1 := &Receiver1{name: "Receiver1"}
	receiver2 := &Receiver2{name: "Receiver2"}

	command1 := &ConcreteCommand{receiver: receiver1}
	command2 := &ConcreteCommand{receiver: receiver2}

	invoker := &Invoker{}
	invoker.SetCommand(command1)
	invoker.ExecuteCommand()

	invoker.SetCommand(command2)
	invoker.ExecuteCommand()
}
