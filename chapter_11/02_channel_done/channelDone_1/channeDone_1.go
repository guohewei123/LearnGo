package main

import (
	"fmt"
)

// 创建 worker 函数
func doWork(i int, ch chan int, done chan bool) {
	for v := range ch{
		fmt.Printf("worker %d received %c\n", i, v)
		done <- true
	}

}

// 创建 worker 并返回 channel
func createWorker(i int) worker {
	c := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(i, c.in, c.done)
	return c
}

type worker struct {
	in chan int
	done chan bool
}


// 通过 sync.WaitGroup 来完成 channel 结束的 demo
func chanDemo() {

	// 创建 worker
	var workers [10] worker
	for i := range workers {
		workers[i] = createWorker(i)
	}

	// 发送数据
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	// 等待 worker 结束
	for _, worker := range workers {
		<- worker.done
	}

	// 发送数据
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	// 等待 worker 结束
	for _, worker := range workers {
		<- worker.done
	}

}


func main() {
	chanDemo()   // 通过两个 channel 双向通信，实现等待worker结束的功能
}
