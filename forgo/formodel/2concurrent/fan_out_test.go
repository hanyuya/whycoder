package formodel

import (
	"fmt"
	"sync"
	"testing"
)

/*
扇出模式（Fan-Out Pattern）是一种并发编程模式，它将一个输入任务分发给多个并发处理器（也称为 worker），以提高处理速度。

	在 Go 语言中，可以使用 goroutine 和 channel 来实现扇出模式。
*/
var wg sync.WaitGroup

func worker(id int, jobs <-chan int, results chan<- int) {
	defer wg.Done()
	for job := range jobs {
		fmt.Println("Worker", id, "processing job", job)
		results <- job * 2
	}
}

func Test_FanOut(t *testing.T) {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs)

	for result := range results {
		fmt.Println("Result:", result)
	}
}
