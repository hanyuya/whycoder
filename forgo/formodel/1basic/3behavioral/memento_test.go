package formodel

import (
	"fmt"
	"testing"
)

/*
备忘录模式（Memento Pattern）

	备忘录模式（Memento Pattern）是一种行为型设计模式，它允许在不暴露对象实现的情况下捕获对象的内部状态并在对象之外保存这个状态，以便稍后可以将其还原到先前的状态。

基本结构

	备忘录模式包括以下几个重要角色：
		发起人Originator： 需要还原状态的那个对象，负责创建一个【备忘录】，并使用备忘录记录当前时刻的内部状态。
		备忘录Memento: 存储发起人对象的内部状态，它可以包含发起人的部分或全部状态信息，但是对外部是不可见的，只有发起人能够访问备忘录对象的状态。
		备忘录有两个接口，发起人能够通过宽接口访问数据，管理者只能看到窄接口，并将备忘录传递给其他对象。
		管理者Caretaker: 负责存储备忘录对象，但并不了解其内部结构，管理者可以存储多个备忘录对象。
		客户端：在需要恢复状态时，客户端可以从管理者那里获取备忘录对象，并将其传递给发起人进行状态的恢复

使用场景
 1. 备忘录模式在保证了对象内部状态的封装和私有性前提下可以轻松地添加新的备忘录和发起人，实现“备份”，不过 备份对象往往会消耗较多的内存，资源消耗增加。
 2. 备忘录模式常常用来实现撤销和重做功能，比如在Java Swing GUI编程中，javax.swing.undo包中的撤销（undo）和重做（redo）机制使用了备忘录模式。
    UndoManager和UndoableEdit接口是与备忘录模式相关的主要类和接口。
*/
type Originator interface {
	Save() Memento
	Restore(memento Memento)
}

type Memento interface {
	GetState() string
}

/* TextEditor 结构体，该结构体实现了 Originator 接口 */
type TextEditor struct {
	state string
}

func (t *TextEditor) Save() Memento {
	return &textMemento{state: t.state}
}

func (t *TextEditor) Restore(memento Memento) {
	t.state = memento.GetState()
}

func (t *TextEditor) SetState(state string) {
	t.state = state
}

func (t *TextEditor) GetState() string {
	return t.state
}

/* textMemento 结构体实现了 Memento 接口，它保存了文本编辑器的状态 */
type textMemento struct {
	state string
}

func (t *textMemento) GetState() string {
	return t.state
}

// Caretaker 结构体负责管理 Memento
type Caretaker struct {
	mementos   []Memento
	currentIdx int
}

func (c *Caretaker) AddMemento(memento Memento) {
	c.mementos = append(c.mementos, memento)
	c.currentIdx = len(c.mementos) - 1
}

func (c *Caretaker) Undo(t *TextEditor) {
	if c.currentIdx > 0 {
		c.currentIdx--
		t.Restore(c.mementos[c.currentIdx])
	}
}

func (c *Caretaker) Redo(t *TextEditor) {
	if c.currentIdx < len(c.mementos)-1 {
		c.currentIdx++
		t.Restore(c.mementos[c.currentIdx])
	}
}

func TestMemento(t *testing.T) {
	editor := TextEditor{}
	caretaker := Caretaker{}

	editor.SetState("State #1")
	caretaker.AddMemento(editor.Save())
	editor.SetState("State #2")
	caretaker.AddMemento(editor.Save())
	editor.SetState("State #3")
	caretaker.AddMemento(editor.Save())

	caretaker.Undo(&editor)
	fmt.Println("Current state:", editor.GetState())

	caretaker.Redo(&editor)
	fmt.Println("Current state:", editor.GetState())
}
