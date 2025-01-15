package formodel

import "sync"

// 单例模式 ———— 创建型设计模式
/*
单例模式是一种创建型设计模式， 它的核心思想是保证一个类只有一个实例，并提供一个全局访问点来访问这个实例。
单例模式的适用场景包括：
一些只允许存在一个实例的类，比如全局统一的监控统计模块
一些实例化时很耗费资源的类，比如协程池、连接池、和第三方交互的客户端、缓存实例等
一些入参繁杂的系统模块组件，比如 controller、service、dao 等

在许多流行的工具和库中，也都使用到了单例设计模式，比如Java中的Runtime类就是一个经典的单例，表示程序的运行时环境。
此外 Spring 框架中的应用上下文 (ApplicationContext) 也被设计为单例，以提供对应用程序中所有 bean 的集中式访问点。
*/

// 饿汉式
/*
singleton 是需要被单例模式保护的类型
singleton 首字母小写，是不可导出的类型，避免被外界直接获取
在包初始化函数 init 中完成了 singleton 单例的初始化工作
对外暴露可导出方法 GetInstance，返回提前初始化好的全局单例对象
*/
type hungerSingleton struct{}

type HungerInstance interface {
	Work()
}

var s *hungerSingleton

func init() {
	s = newHungerSingleton()
}

func newHungerSingleton() *hungerSingleton {
	return &hungerSingleton{}
}

func (s *hungerSingleton) Work() {}

func GetHungerInstance() HungerInstance {
	return s
}

//---------------------------------------

// 懒汉式--double check lock
type lazySingleton struct{}

type LazyInstance interface {
	LazyWork()
}

var (
	ls    *lazySingleton // 指针零值为nil
	mutex sync.Mutex     // 可以直接使用 零值初始化
	once  sync.Once
)

func newLazySingleton() *lazySingleton {
	return &lazySingleton{}
}

func (s *lazySingleton) LazyWork() {}

func GetLazyInstance() LazyInstance {

	if ls != nil {
		return ls
	}
	mutex.Lock()
	defer mutex.Unlock()
	if ls != nil {
		return ls
	}
	ls = newLazySingleton()
	return ls
}

// 懒汉式--sync.once
func GetLazyInstance2() LazyInstance {
	once.Do(func() {
		ls = newLazySingleton()
	})
	return ls
}
