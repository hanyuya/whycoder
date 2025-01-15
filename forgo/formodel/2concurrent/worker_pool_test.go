package formodel

import (
	"fmt"
	"sync"
	"testing"
)

/*
工作池模式 (Worker Pool Pattern) 是一种并发模式，用于通过限制并发执行的任务数量来提高应用程序的性能和可伸缩性。
	该模式涉及创建一个工作池，其中包含固定数量的工作者（Worker），这些工作者从一个共享的队列中获取任务，并执行它们。
	当一个工作者完成任务时，它会返回工作池，等待分配新的任务。
*/

type Task struct {
	ID int
}

func poolWorker(id int, taskQueue chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range taskQueue {
		fmt.Printf("Worker %d processing task %d\n", id, task.ID)
	}
}

func Test_pool_worker(t *testing.T) {
	taskQueue := make(chan Task, 10)
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go poolWorker(i, taskQueue, wg)
	}

	for i := 0; i < 10; i++ {
		taskQueue <- Task{ID: i}
	}
	close(taskQueue)
	wg.Wait()
}
