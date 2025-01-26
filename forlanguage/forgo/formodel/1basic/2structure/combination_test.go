package formodel

import (
	"fmt"
	"testing"
)

/*
组合模式
是一种结构型设计模式，它将对象组合成树状结构来表示“部分-整体”的层次关系。
组合模式使得客户端可以统一处理单个对象和对象的组合，而无需区分它们的具体类型。

组合模式包括下面几个角色：
	理解起来比较抽象，我们用“省份-城市”举个例子，省份中包含了多个城市，如果将之比喻成一个树形结构，
	城市就是叶子节点，它是省份的组成部分，而“省份”就是合成节点，可以包含其他城市，形成一个整体，省份和城市都是组件，它们都有一个共同的操作，比如获取信息。

Component组件： 组合模式的“根节点”，定义组合中所有对象的通用接口，可以是抽象类或接口。该类中定义了子类的共性内容。
Leaf叶子：实现了Component接口的叶子节点，表示组合中的叶子对象，叶子节点没有子节点。
Composite合成： 作用是存储子部件，并且在Composite中实现了对子部件的相关操作，比如添加、删除、获取子组件等。
通过组合模式，整个省份的获取信息操作可以一次性地执行，而无需关心省份中的具体城市。这样就实现了对国家省份和城市的管理和操作。

使用场景
组合模式可以使得客户端可以统一处理单个对象和组合对象，无需区分它们之间的差异，
	比如在图形编辑器中，图形对象可以是简单的线、圆形，也可以是复杂的组合图形，这个时候可以对组合节点添加统一的操作。
	总的来说，组合模式适用于任何需要构建具有部分-整体层次结构的场景，比如组织架构管理、文件系统的文件和文件夹组织等。
*/

type CComponent interface {
	Operation()
}

type Leaf struct {
	name string
}

func (l *Leaf) Operation() {
	fmt.Println("leaf operation")
}

type Composite struct {
	name     string
	children []CComponent
}

func NewComposite(name string) *Composite {
	return &Composite{
		name:     name,
		children: make([]CComponent, 0),
	}
}

func (c *Composite) add(cc CComponent) {
	c.children = append(c.children, cc)
}

func (c *Composite) remove() {
}

func (c *Composite) Operation() {
	for _, c := range c.children {
		c.Operation()
	}
}

func TestComposite(t *testing.T) {
	myComposite := NewComposite("new composite")
	leaf1 := &Leaf{name: "leaf1"}
	leaf2 := &Leaf{name: "leaf2"}
	myComposite.add(leaf1)
	myComposite.add(leaf2)
	leaf1.Operation()
	myComposite.Operation()
}
