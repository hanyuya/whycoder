package formodel

/*
抽象工厂模式包含多个抽象产品接口，多个具体产品类，一个抽象工厂接口和多个具体工厂，每个具体工厂负责创建一组相关的产品。

抽象产品接口AbstractProduct: 定义产品的接口，可以定义多个抽象产品接口，比如说沙发、椅子、茶几都是抽象产品。
具体产品类ConcreteProduct: 实现抽象产品接口，产品的具体实现，古典风格和沙发和现代风格的沙发都是具体产品。
抽象工厂接口AbstractFactory: 声明一组用于创建产品的方法，每个方法对应一个产品。
具体工厂类ConcreteFactory： 实现抽象工厂接口，负责创建一组具体产品的对象，在本例中，生产古典风格的工厂和生产现代风格的工厂都是具体实例。


应用场景
抽象工厂模式能够保证一系列相关的产品一起使用，并且在不修改客户端代码的情况下，可以方便地替换整个产品系列。
但是当需要增加新的产品类时，除了要增加新的具体产品类，还需要修改抽象工厂接口及其所有的具体工厂类，扩展性相对较差。
因此抽象工厂模式特别适用于一系列相关或相互依赖的产品被一起创建的情况，典型的应用场景是使用抽象工厂模式来创建与不同数据库的连接对象。

简单工厂、工厂方法、抽象工厂的区别:
简单工厂模式：一个工厂方法创建所有具体产品
工厂方法模式：一个工厂方法创建一个具体产品
抽象工厂模式：一个工厂方法可以创建一类具体产品
*/

type ProductA interface {
	Display()
}
type ProductB interface {
	Show()
}

type ProductA1 struct{}

func (p *ProductA1) Display() {}

type ProductA2 struct{}

func (p *ProductA2) Display() {}

type ProductB1 struct{}

func (p *ProductB1) Show() {}

type ProductB2 struct{}

func (p *ProductB2) Show() {}

type AbstractFactory interface {
	CreateProductA() ProductA
	CreateProductB() ProductB
}

type ConcreteFactory1 struct{}

func (c *ConcreteFactory1) CreateProductA() ProductA {
	return &ProductA1{}
}

func (c *ConcreteFactory1) CreateProductB() ProductB {
	return &ProductB1{}
}

type ConcreteFactory2 struct{}

func (c *ConcreteFactory2) CreateProductA() ProductA {
	return &ProductA2{}
}

func (c *ConcreteFactory2) CreateProductB() ProductB {
	return &ProductB2{}
}

func test_abstractFactory() {
	cf1 := &ConcreteFactory1{}
	pa1 := cf1.CreateProductA()
	pb1 := cf1.CreateProductB()
	pa1.Display()
	pb1.Show()

	cf2 := &ConcreteFactory2{}
	pa2 := cf2.CreateProductA()
	pb2 := cf2.CreateProductB()
	pa2.Display()
	pb2.Show()
}
