package formodel

import (
	"fmt"
	"testing"
)

/*
迭代器模式是一种行为设计模式，是一种使用频率非常高的设计模式，在各个语言中都有应用，
	其主要目的是提供一种统一的方式来访问一个聚合对象中的各个元素，而不需要暴露该对象的内部表示。
	通过迭代器，客户端可以顺序访问聚合对象的元素，而无需了解底层数据结构。

迭代器模式应用广泛，但是大多数语言都已经内置了迭代器接口，不需要自己实现。

迭代器模式包括以下几个重要角色
1. 迭代器接口Iterator：定义访问和遍历元素的接口, 通常会包括：
	- hasNext()方法用于检查是否还有下一个元素
	- next()方法用于获取下一个元素。
	- 获取第一个元素
	- 获取当前元素的方法。
2. 具体迭代器ConcreateIterator：实现迭代器接口，实现遍历逻辑对聚合对象进行遍历。
3. 抽象聚合类：定义了创建迭代器的接口，包括一个createIterator方法用于创建一个迭代器对象。
4. 具体聚合类：实现在抽象聚合类中声明的createIterator() 方法，返回一个与具体聚合对应的具体迭代器

使用场景
1. 迭代器模式是一种通用的设计模式，其封装性强，简化了客户端代码，客户端不需要知道集合的内部结构，只需要关心迭代器和迭代接口就可以完成元素的访问。
	但是引入迭代器模式会增加额外的类，每增加一个集合类，都需要增加该集合对应的迭代器，这也会使得代码结构变得更加复杂。

2. 许多编程语言和框架都使用了这个模式提供一致的遍历和访问集合元素的机制。下面是几种常见语言迭代器模式的实现。

Java语言：集合类（如ArrayList、LinkedList), 通过Iterator接口，可以遍历集合中的元素。
	List<String> list = new ArrayList<>();
	list.add("Item 1");
	list.add("Item 2");
	list.add("Item 3");

	Iterator<String> iterator = list.iterator();
	while (iterator.hasNext()) {
		System.out.println(iterator.next());
	}

Python语言：使用迭代器和生成器来实现迭代模式，iter()和next()函数可以用于创建和访问迭代器。
	elements = ["Element 1", "Element 2", "Element 3"]
	iterator = iter(elements)

	while True:
		try:
			element = next(iterator)
			print(element)
		except StopIteration:
			break

C++语言：C++中的STL提供了迭代器的支持，begin()和end()函数可以用于获取容器的起始和结束迭代器。
	#include <iostream>
	#include <vector>

	int main() {
		std::vector<std::string> elements = {"Element 1", "Element 2", "Element 3"};

		for (auto it = elements.begin(); it != elements.end(); ++it) {
			std::cout << *it << std::endl;
		}

		return 0;
	}

JavaScript语言：ES6中新增了迭代器协议，使得遍历和访问集合元素变得更加方便。
	// 可迭代对象实现可迭代协议
	class IterableObject {
	constructor() {
		this.elements = [];
	}
	addElement(element) {
		this.elements.push(element);
	}
	[Symbol.iterator]() {
		let index = 0;
		// 迭代器对象实现迭代器协议
		return {
		next: () => {
			if (index < this.elements.length) {
			return { value: this.elements[index++], done: false };
			} else {
			return { done: true };
			}
		}
		};
	}
	}

	// 使用迭代器遍历可迭代对象
	const iterableObject = new IterableObject();
	iterableObject.addElement("Element 1");
	iterableObject.addElement("Element 2");
	iterableObject.addElement("Element 3");

	for (const element of iterableObject) {
	console.log(element);
	}
*/

// Iterator 接口
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// Aggregate 聚合对象接口
type Aggregate interface {
	Iterator() Iterator
}

// ConcreteAggregate 具体聚合对象类型
type NumberAggregate struct {
	numbers []int
}

func (na *NumberAggregate) Iterator() Iterator {
	return &NumberIterator{
		numbers: na.numbers,
		index:   0,
	}
}

// ConcreteIterator 具体迭代器类型
type NumberIterator struct {
	numbers []int
	index   int
}

func (ni *NumberIterator) HasNext() bool {
	return ni.index < len(ni.numbers)
}

func (ni *NumberIterator) Next() interface{} {
	number := ni.numbers[ni.index]
	ni.index++
	return number
}

// TestIterator 测试迭代器模式
func TestIterator(t *testing.T) {
	na := &NumberAggregate{
		numbers: []int{1, 2, 3, 4, 5},
	}

	it := na.Iterator()
	for it.HasNext() {
		fmt.Println(it.Next())
	}
}
