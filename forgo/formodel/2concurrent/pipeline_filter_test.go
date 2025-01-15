package formodel

import (
	"fmt"
	"strings"
	"testing"
)

/*
管道和过滤器模式是一种常见的软件设计模式，它可以将一个大的处理任务拆分成多个小的处理步骤，每个步骤都是一个独立的过滤器，它们串联起来形成一个管道。
	每个过滤器只负责一个简单的任务，比如转换数据格式、过滤掉无用信息等。
	通过这种方式，可以使得复杂的处理任务更加可控、易于维护。

在 Go 语言中，可以使用管道和过滤器模式来处理一些数据处理任务，比如日志处理、图像处理等。
	管道是一种可以实现并发处理的数据结构，它可以将数据从一个过滤器传递到另一个过滤器，并最终输出处理结果。
	过滤器是实际进行数据处理的组件，它们可以被串联起来形成一个处理管道。

下面是一个使用管道和过滤器模式实现简单的字符串处理任务的示例代码。代码包含了三个过滤器：
	1、source 过滤器：生成一些字符串数据，并将它们输出到管道中。
	2、uppercase 过滤器：将输入的字符串转换成大写字母，并将转换后的字符串输出到管道中。
	3、sink 过滤器：从管道中读取数据，并将其输出到标准输出中。

值得注意的是，当 source 过滤器输出完所有数据后，它会关闭输出管道。而当 uppercase 过滤器处理完所有数据后，它也会关闭输出管道。
	这样，在管道中的所有数据被处理完后，最终的 sink 过滤器会从输入管道中读取到一个关闭信号，从而结束处理过程。

使用管道和过滤器模式可以轻松地将一个大的处理任务拆分成多个小的处理步骤，并实现并发处理。
	在 Go 语言中，管道和过滤器模式是一种非常常用的处理数据的方式，它能够有效地提高数据处理的效率和可维护性。

*/

func source(out chan<- string) {
	// 产生一些字符串数据并将它们输出到管道中
	data := []string{"hello", "world", "go", "language"}
	for _, s := range data {
		out <- s
	}
	close(out)
}

func uppercase(in <-chan string, out chan<- string) {
	// 从管道中读取数据，并将其转换成大写字母，并将转换后的字符串输出到管道中
	for s := range in {
		out <- strings.ToUpper(s)
	}
	close(out)
}

func sink(in <-chan string, end chan struct{}) {
	// 从管道中读取数据，并将其输出到标准输出中
	for s := range in {
		fmt.Println(s)
	}
	close(end)
}

func Test_pipeline(t *testing.T) {
	sourceChan := make(chan string)
	uppercaseChan := make(chan string)
	sinkChan := make(chan struct{})

	// 创建并启动三个过滤器
	go source(sourceChan)
	go uppercase(sourceChan, uppercaseChan)
	go sink(uppercaseChan, sinkChan)

	// 等待三个过滤器完成处理
	<-sinkChan
}
