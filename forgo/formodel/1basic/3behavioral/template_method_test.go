package formodel

import (
	"fmt"
	"testing"
)

/*
模板方法模式（Template Method Pattern）是一种行为型设计模式,
	它定义了一个算法的骨架，将一些步骤的实现延迟到子类。
	模板方法模式使得子类可以在不改变算法结构的情况下，重新定义算法中的某些步骤。

模板方法模式的基本结构包含以下两个角色：
1. 模板类AbstractClass：由一个模板方法和若干个基本方法构成，模板方法定义了逻辑的骨架，按照顺序调用包含的基本方法，
		基本方法通常是一些抽象方法，这些方法由子类去实现。基本方法还包含一些具体方法，它们是算法的一部分但已经有默认实现，在具体子类中可以继承或者重写。
2. 具体类ConcreteClass：继承自模板类，实现了在模板类中定义的抽象方法，以完成算法中特定步骤的具体实现。

应用场景
1. 模板方法模式将算法的不变部分被封装在模板方法中，而可变部分算法由子类继承实现，这样做可以很好的提高代码的复用性，
	但是当算法的框架发生变化时，可能需要修改模板类，这也会影响到所有的子类。

2. 总体来说，当算法的整体步骤很固定，但是个别步骤在更详细的层次上的实现可能不同时，通常考虑模板方法模式来处理。

3. 在已有的工具和库中， Spring框架中的JdbcTemplate类使用了模板方法模式，其中定义了一些执行数据库操作的模板方法，具体的数据库操作由回调函数提供。
	而在Java的JDK源码中，AbstractList 类也使用了模板方法模式，它提供了一些通用的方法，其中包括一些模板方法。具体的列表操作由子类实现。

*/

// 定义一个抽象类
type AbstractClass interface {
	// 模板方法
	TemplateMethod()
	// 基本方法1
	BaseMethod1()
	// 基本方法2
	BaseMethod2()
}

// 定义一个具体类1
type ConcreteClass struct {
	AbstractClass
}

// 实现模板方法
func (c *ConcreteClass) TemplateMethod() {
	fmt.Println("模板方法: 算法的骨架")
	fmt.Println("基本方法1: 步骤1")
	c.BaseMethod1()
	fmt.Println("基本方法2: 步骤2")
	c.BaseMethod2()
}

// 实现基本方法1
func (c *ConcreteClass) BaseMethod1() {
	// 具体实现
	fmt.Println("具体类1基本方法1: 步骤1的具体实现")
}

// 实现基本方法2
func (c *ConcreteClass) BaseMethod2() {
	// 具体实现
	fmt.Println("具体类1基本方法2: 步骤2的具体实现")
}

// 定义一个具体类2
type ConcreteClass2 struct {
	AbstractClass
}

// 实现模板方法
func (c *ConcreteClass2) TemplateMethod() {
	fmt.Println("模板方法: 算法的骨架")
	fmt.Println("基本方法1: 步骤1")
	c.BaseMethod1()
	fmt.Println("基本方法2: 步骤2")
	c.BaseMethod2()
}

// 实现基本方法1
func (c *ConcreteClass2) BaseMethod1() {
	// 具体实现
	fmt.Println("具体类2基本方法1: 步骤1的具体实现")
}

// 实现基本方法2
func (c *ConcreteClass2) BaseMethod2() {
	// 具体实现
	fmt.Println("具体类2基本方法2: 步骤2的具体实现")
}

// 测试
func TestTemplateMethod(t *testing.T) {
	// 实例化具体类1
	c1 := ConcreteClass{}
	// 调用模板方法
	c1.TemplateMethod()
	fmt.Println("------------------------------------------")
	// 实例化具体类2
	c2 := ConcreteClass2{}
	// 调用模板方法
	c2.TemplateMethod()
}
