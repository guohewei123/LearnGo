package main

import (
	"fmt"
	"sync"
)

// 创建 worker 函数
func doWork(i int, w worker) {
	for v := range w.in {
		fmt.Printf("worker %d received %c\n", i, v)
		w.done()
	}
}

// 创建 worker 并返回 channel
func createWorker(i int, wg *sync.WaitGroup) worker {
	w := worker{
		in:   make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(i, w)
	return w
}

// 创建 worker 结构体
type worker struct {
	in   chan int // 发送数据 channel
	done func()   // 结束函数
}

// 通过 sync.WaitGroup 实现等待 worker 结束的功能
func chanWaitGroupDemo() {

	// 创建 worker
	var wg sync.WaitGroup  // 创建 WaitGroup
	var workers [10]worker
	for i := range workers {
		workers[i] = createWorker(i, &wg)
	}

	//wg.Add(20)           // 添加 20 个 WaitGroup, 代码可以一次性添加，也可以执行一个work 添加一个
	// 发送数据
	for i, worker := range workers {
		worker.in <- 'a' + i
		wg.Add(1)          // 添加 20 个 WaitGroup, 代码可以一次性添加，也可以执行一个work 添加一个
	}
	// 发送数据
	for i, worker := range workers {
		worker.in <- 'A' + i
		wg.Add(1)          // 添加 20 个 WaitGroup, 代码可以一次性添加，也可以执行一个work 添加一个
	}
	// 等待 worker 结束
	wg.Wait()

}

func main() {
	chanWaitGroupDemo() // 通过 sync.WaitGroup 实现等待 worker 结束的功能
}
