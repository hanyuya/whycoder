package formodel

import "fmt"

/*
建造者模式（也被成为生成器模式），是一种创建型设计模式，软件开发过程中有的时候需要创建很复杂的对象，
而建造者模式的主要思想是:
将对象的构建过程分为多个步骤，并为每个步骤定义一个抽象的接口。
具体的构建过程由实现了这些接口的具体建造者类来完成。
同时有一个指导者类负责协调建造者的工作，按照一定的顺序或逻辑来执行构建步骤，最终生成产品。

建造者模式有下面几个关键角色：
产品Product：被构建的复杂对象, 包含多个组成部分。
抽象建造者Builder: 定义构建产品各个部分的抽象接口和一个返回复杂产品的方法getResult
具体建造者Concrete Builder：实现抽象建造者接口，构建产品的各个组成部分，并提供一个方法返回最终的产品。
指导者Director：调用具体建造者的方法，按照一定的顺序或逻辑来构建产品。
*/

type ProductBu struct {
	Part1 string
	Part2 string
}

func (p *ProductBu) SetPart1(part string) {
	p.Part1 = part
}

func (p *ProductBu) SetPart2(part string) {
	p.Part2 = part
}

type Builder interface {
	BuildPart1(part string)
	BuildPart2(part string)
	GetResult() *ProductBu
}

type ConcreteBuilder struct {
	ProductInstance *ProductBu
}

func (cb *ConcreteBuilder) BuildPart1(part string) {
	cb.ProductInstance.Part1 = part
}

func (cb *ConcreteBuilder) BuildPart2(part string) {
	cb.ProductInstance.Part2 = part
}

func (cb *ConcreteBuilder) GetResult() *ProductBu {
	return cb.ProductInstance
}

type Director struct {
	BuilderInstance Builder
}

func (d *Director) Construct() {
	d.BuilderInstance.BuildPart1("part1")
	d.BuilderInstance.BuildPart2("part1")
}

func TestBuilder() {
	bi := ConcreteBuilder{}
	director := Director{
		BuilderInstance: &bi,
	}
	director.Construct()

	product := bi.GetResult()

	fmt.Printf("product.Part1: %v\n", product.Part1)
	fmt.Printf("product.Part2: %v\n", product.Part2)
}
