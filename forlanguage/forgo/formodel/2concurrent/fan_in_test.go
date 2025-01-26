package formodel

import (
	"fmt"
	"sync"
	"testing"
)

/*
扇入模式（Fan-In Pattern）是一种并发编程模式，用于将多个输入通道（channels）中的数据聚合成一个输出通道。
	具体来说，扇入模式通过从每个输入通道中读取数据，并将它们合并到一个输出通道中，从而实现了数据的聚合。
	扇入模式通常用于并发任务中，可以显著提高程序的吞吐量和性能。
*/

var done = make(chan struct{})

func fanInProducer(ch chan<- int, num int) {
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
}

func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(channels))
	for _, ch := range channels {
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				select {
				case out <- v:
					fmt.Println("in out")
				case <-done:
					fmt.Println("done")
					return
				}
			}
			fmt.Println("ch is closed")
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func Test_FanIn(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go fanInProducer(ch1, 1)
	go fanInProducer(ch2, 2)
	go fanInProducer(ch3, 3)
	for v := range fanIn(ch1, ch2, ch3) {
		fmt.Println("v:", v)
	}
	close(done)
	fmt.Println("main done")
}
