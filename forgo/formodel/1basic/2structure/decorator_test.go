package formodel

import (
	"fmt"
	"testing"
)

/*
装饰模式（Decorator Pattern, 结构型设计模式）
	通常情况下，扩展类的功能可以通过继承实现，但是扩展越多，子类越多
	装饰模式（Decorator Pattern, 结构型设计模式）可以在不定义子类的情况下动态的给对象添加一些额外的功能。
	具体的做法是将原始对象放入包含行为的特殊封装类(装饰类)，从而为原始对象动态添加新的行为，而无需修改其代码。

	举个简单的例子，假设你有一个基础的图形类，你想要为图形类添加颜色、边框、阴影等功能，如果每个功能都实现一个子类，就会导致产生大量的类，
	这时就可以考虑使用装饰模式来动态地添加，而不需要修改图形类本身的代码，这样可以使得代码更加灵活、更容易维护和扩展。

基本结构：
装饰模式包含以下四个主要角色：
	组件Component：通常是抽象类或者接口，是具体组件和装饰者的父类，定义了具体组件需要实现的方法，比如说我们定义Coffee为组件。
	具体组件ConcreteComponent: 实现了Component接口的具体类，是被装饰的对象。
	装饰类Decorator: 一个抽象类，给具体组件添加功能，但是具体的功能由其子类具体装饰者完成，持有一个指向Component对象的引用。
	具体装饰类ConcreteDecorator: 扩展Decorator类，负责向Component对象添加新的行为，加牛奶的咖啡是一个具体装饰类，加糖的咖啡也是一个具体装饰类。

应用场景
装饰模式通常在以下几种情况使用：
	当需要给一个现有类添加附加功能，但由于某些原因不能使用继承来生成子类进行扩充时，可以使用装饰模式。
	动态的添加和覆盖功能：当对象的功能要求可以动态地添加，也可以再动态地撤销时可以使用装饰模式。
	在Java的I/O库中，装饰者模式被广泛用于增强I/O流的功能。
		例如，BufferedInputStream和BufferedOutputStream这两个类提供了缓冲区的支持，
		通过在底层的输入流和输出流上添加缓冲区，提高了读写的效率，它们都是InputStream和OutputStream的装饰器。
		BufferedReader和BufferedWriter这两个类与BufferedInputStream和BufferedOutputStream类似，提供了字符流的缓冲功能，是Reader和Writer的装饰者。
*/

type Component interface {
	Operation()
}

type ConcreteComponent struct{}

func (cc *ConcreteComponent) Operation() {
	fmt.Println("concrete component operation")
}

// type AbstractDecorator struct {
// 	component Component
// }

// func (ad *AbstractDecorator) Operation() {
// 	ad.component.Operation()
// }

type AbstractDecorator Component

type ConcreteDecorator struct {
	AbstractDecorator
}

func (cd *ConcreteDecorator) Operation() {
	fmt.Println("other func prefix")
	cd.AbstractDecorator.Operation()
	fmt.Println("other func suffix")
}

func TestDecorator(t *testing.T) {
	cc := &ConcreteComponent{}
	cd := ConcreteDecorator{
		AbstractDecorator: cc,
	}
	cd.Operation()
}
