package formodel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
生产者消费者模式 (Producer-Consumer Pattern)
	是一种并发编程模式，用于解决生产者和消费者之间数据交换的问题。
	在这种模式中，生产者生产数据，并将其传递给消费者进行处理。
	生产者和消费者是独立的进程或线程，它们通过共享的缓冲区进行通信。

在 Go 语言中，可以使用通道 (Channel) 来实现生产者消费者模式。
	通道是一种用于在 Go 协程之间传递数据的特殊数据类型，它具有阻塞式的特性，即当通道已满或者为空时，写入或读取数据的操作会被阻塞。
*/

/*
	n : 1

多个生产者对应一个消费者，即 N:1
消费者处理生产者发送过来的消息时是并行处理的，但是有速率限制，最大为5qps
消费者处理完了某个生产者的消息后，通知对应的生产者
当某个生产者发送的所有消息都收到处理完成的消息后，执行后续逻辑
*/
const rateLimit = 5

type msgChan struct {
	Id       int64
	Text     string
	DoneChan chan int64
}

var msgCh chan msgChan

func initChan() {
	msgCh = make(chan msgChan, 10)
}

// func SendToChan(msg msgChan) {
// 	msgCh <- msg
// }

// func GetFromChan() chan msgChan {
// 	return msgCh
// }

func producer(i int64) *msgChan {
	var msg = msgChan{
		Id:       i,
		Text:     fmt.Sprintf("消息文本%v", i),
		DoneChan: make(chan int64),
	}
	msgCh <- msg
	fmt.Println("producer: ", i)
	return &msg
}

func consumer() {

	// for msg := range GetFromChan() {
	for msg := range msgCh {
		fmt.Printf("consumer %v processing ..., time: %v \n", msg.Id, time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep((1000 / rateLimit) * time.Millisecond)
		msg.DoneChan <- msg.Id
	}

	// for {
	// 	select {
	// 	case msg, ok := <-GetFromChan():
	// 		if ok {
	// 			fmt.Printf("consumer %v processing ..., time: %v\n", msg.Id, time.Now().Format("2006-01-02 15:04:05"))
	// 			time.Sleep((1000 / rateLimit) * time.Millisecond)
	// 			msg.DoneChan <- msg.Id
	// 		}
	// 	}
	// }
}

func TestPC(t *testing.T) {
	initChan()
	go consumer()
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 20; i++ {
		msg := producer(int64(i))
		go func(m *msgChan, w *sync.WaitGroup) {
			defer w.Done()
			if v, ok := <-m.DoneChan; ok {
				fmt.Println("receive done: ", v)
				close(m.DoneChan)
			}
		}(msg, &wg)
	}
	wg.Wait()
	close(msgCh)
	fmt.Println("后续操作")
}
