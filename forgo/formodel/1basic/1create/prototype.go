package formodel

/*
原型模式一种创建型设计模式，该模式的核心思想是基于现有的对象创建新的对象，而不是从头开始创建。
在原型模式中，通常有一个原型对象，它被用作创建新对象的模板。新对象通过复制原型对象的属性和状态来创建，而无需知道具体的创建细节。

原型模式包含两个重点模块：
抽象原型接口prototype: 声明一个克隆自身的方法clone
具体原型类ConcretePrototype: 实现clone方法，复制当前对象并返回一个新对象。
*/
type Prototype interface {
	Clone() Prototype
}

type ConcretePrototype struct {
	data string
}

func (cp *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{
		data: cp.data,
	}
}

func (cp *ConcretePrototype) getData() string {
	return cp.data
}

func Test_Clone() {
	var p Prototype
	p = &ConcretePrototype{"i am data"}

	p2 := p.Clone()

	if cp, ok := p2.(*ConcretePrototype); ok {
		println(cp.getData())
	}
}
