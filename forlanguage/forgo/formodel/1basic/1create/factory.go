package formodel

/*
简单工厂：
一个具体工厂、抽象产品、具体产品
*/

/*
工厂方法模式分为以下几个角色：

抽象工厂：一个接口，包含一个抽象的工厂方法（用于创建产品对象）。
具体工厂：实现抽象工厂接口，创建具体的产品。
抽象产品：定义产品的接口。
具体产品：实现抽象产品接口，是工厂创建的对象。

应用场景：
工厂方法模式使得每个工厂类的职责单一，每个工厂只负责创建一种产品，
当创建对象涉及一系列复杂的初始化逻辑，而这些逻辑在不同的子类中可能有所不同时，
可以使用工厂方法模式将这些初始化逻辑封装在子类的工厂中。在现有的工具、库中，工厂方法模式也有广泛的应用，比如：

Spring 框架中的 Bean 工厂：通过配置文件或注解，Spring 可以根据配置信息动态地创建和管理对象。
JDBC 中的 Connection 工厂：在 Java 数据库连接中，DriverManager 使用工厂方法模式来创建数据库连接。不同的数据库驱动（如 MySQL、PostgreSQL 等）都有对应的工厂来创建连接。
*/
type Shape interface {
	Draw()
}

type Circle struct{}

func (c *Circle) Draw() {}

type Square struct{}

func (s *Square) Draw() {}

type ShapeFactory interface {
	CreateShape() Shape
}

type CircleFactory struct{}

func (c *CircleFactory) CreateShape() Shape {
	return &Circle{}
}

type SquareFactory struct{}

func (c *SquareFactory) CreateShape() Shape {
	return &Square{}
}

func testFactory() {

	circleFactory := &CircleFactory{}
	circle := circleFactory.CreateShape()
	circle.Draw()

	squareFactory := &SquareFactory{}
	square := squareFactory.CreateShape()
	square.Draw()
}
