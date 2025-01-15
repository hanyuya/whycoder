package formodel

import (
	"fmt"
	"testing"
)

/*
访问者模式（Visitor Pattern）是一种行为型设计模式，可以在不改变对象结构的前提下，对对象中的元素进行新的操作。

举个例子，假设有一个动物园，里面有不同种类的动物，比如狮子、大象、猴子等。每个动物都会被医生检查身体，被管理员投喂，被游客观看，医生，游客，管理员都属于访问者。
如果你想对动物园中的每个动物执行一些操作，比如医生健康检查、管理员喂食、游客观赏等。就可以使用访问者模式来实现这些操作。
将这些访问者应用到动物园的每个动物上。

基本结构：
访问者模式包括以下几个基本角色：
1. 抽象访问者（Visitor）： 声明了访问者可以访问哪些元素，以及如何访问它们的方法visit。
2. 具体访问者（ConcreteVisitor）： 实现了抽象访问者定义的方法，不同的元素类型可能有不同的访问行为。医生、管理员、游客都属于具体的访问者，它们的访问行为不同。
3. 抽象元素（Element）： 定义了一个accept方法，用于接受访问者的访问。
4. 具体元素（ConcreteElement）： 实现了accept方法，是访问者访问的目标。
5. 对象结构（Object Structure）： 包含元素的集合，可以是一个列表、一个集合或者其他数据结构。负责遍历元素，并调用元素的接受方法

使用场景
访问者模式结构较为复杂，但是访问者模式将同一类操作封装在一个访问者中，使得相关的操作彼此集中，提高了代码的可读性和维护性。
它常用于对象结构比较稳定，但经常需要在此对象结构上定义新的操作，这样就无需修改现有的元素类，只需要定义新的访问者来添加新的操作。
*/

type Visitor interface {
	VisitElementA(element *ElementA)
	VisitElementB(element *ElementB)
}

type Element interface {
	Accept(visitor Visitor)
}

type ElementA struct {
	name string
}

func (e *ElementA) Accept(visitor Visitor) {
	visitor.VisitElementA(e)
}

type ElementB struct {
	name string
}

func (e *ElementB) Accept(visitor Visitor) {
	visitor.VisitElementB(e)
}

type ConcreteVisitor struct {
	Name string
}

func (cv *ConcreteVisitor) VisitElementA(element *ElementA) {
	fmt.Println("ConcreteVisitor visiting ElementA:", element.name)
}

func (cv *ConcreteVisitor) VisitElementB(element *ElementB) {
	fmt.Println("ConcreteVisitor visiting ElementB:", element.name)
}

type ObjectStructure struct {
	elements []Element
}

func (os *ObjectStructure) AddElement(element Element) {
	os.elements = append(os.elements, element)
}

func (os *ObjectStructure) Accept(visitor Visitor) {
	for _, element := range os.elements {
		element.Accept(visitor)
	}
}

func TestVisitor(t *testing.T) {
	os := ObjectStructure{}
	os.AddElement(&ElementA{name: "ElementA1"})
	os.AddElement(&ElementB{name: "ElementB1"})
	os.AddElement(&ElementA{name: "ElementA2"})
	os.AddElement(&ElementB{name: "ElementB2"})

	visitor := &ConcreteVisitor{Name: "Visitor1"}
	os.Accept(visitor)
}
