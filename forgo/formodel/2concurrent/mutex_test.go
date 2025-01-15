package formodel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
Go 语言中的 Mutex 模式（互斥锁模式）是一种常见的同步模式，它使用互斥锁来保护共享资源，以防止并发访问导致数据竞争和不一致性。

在 Go 语言中，当多个 goroutine 并发访问共享资源时，可能会发生数据竞争，从而导致程序出现不可预测的行为。
	为了避免这种情况，我们需要使用同步原语来保护共享资源。

互斥锁模式（Mutex Pattern）是一种常见的同步模式，它使用互斥锁（mutex）来实现共享资源的互斥访问。
	当一个 goroutine 需要访问共享资源时，它必须先获取互斥锁，这样其他 goroutine 就无法访问该资源。
	一旦该 goroutine 完成了对共享资源的操作，它就会释放互斥锁，以便其他 goroutine 可以继续访问共享资源。

互斥锁模式可以确保共享资源在任何时候只被一个 goroutine 访问，从而避免了数据竞争和不一致性

总结：
1、互斥锁模式是一种常见的同步模式，它使用互斥锁来保护共享资源，以防止并发访问导致数据竞争和不一致性。
	在 Go 语言中，使用标准库中的 sync 包提供的 Mutex 类型来实现互斥锁模式。
	在访问共享资源时，每个 goroutine 必须先获取互斥锁，以防止其他 goroutine 并发访问。
	一旦 goroutine 完成了对共享资源的操作，它就会释放互斥锁，以便其他 goroutine 可以继续访问共享资源。

2、互斥锁模式可以应用于任何需要保护共享资源的场景，例如对共享变量、共享数据结构、文件、网络连接等外部资源的并发访问。
	使用互斥锁模式可以确保共享资源在任何时候只被一个 goroutine 访问，从而避免了数据竞争和不一致性的问题。

*/

var m = make(map[int]int)
var mutex sync.Mutex

func put(k, v int) {
	mutex.Lock()
	defer mutex.Unlock()
	m[k] = v
}

func Test_Mutex(t *testing.T) {
	for i := 0; i < 10; i++ {
		go put(i, i+1)
	}

	time.Sleep(3 * time.Second)

	for k, v := range m {
		fmt.Println("key:", k, "value:", v)
	}
}
