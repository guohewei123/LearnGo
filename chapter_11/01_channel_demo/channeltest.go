package main

import (
	"fmt"
	"time"
)

// 创建 worker 函数
func worker(i int, ch chan int) {

	/* // 方式一 打印 channel 中收到的数据
	for {
		fmt.Printf("worker %d received %c\n", i, <-ch)
	}
	*/

	/*
	// 方式二 当 channel 关闭的时候，停止接收
	for {
		v, ok := <- ch
		if !ok {
			break
		}
		fmt.Printf("worker %d received %c\n", i, v)
	}
	*/

	// 方式三 当 channel 关闭的时候，停止接收
	for v := range ch{
		fmt.Printf("worker %d received %c\n", i, v)
	}

}

// 创建 worker 并返回 channel
func createWorker(i int) chan int {
	ch := make(chan int)
	go worker(i, ch)
	return ch
}

// 返回只允许发送的channel
func createWorkerReturnSendChan(i int) chan<- int {
	ch := make(chan int, 5)
	go worker(i, ch)
	return ch
}

// channel 使用 demo
func chanDemo() {
	//var channels [10]chan int
	var channels [10]chan<- int  // 只允许发送的channel
	for i := range channels {
		//channels[i] = createWorker(i)            
		channels[i] = createWorkerReturnSendChan(i) 
	}

	for i := range channels {
		channels[i] <- 'a' + i
		//c := <- channels[i]  //  invalid operation: <-channels[i] (receive from send-only type chan<- int)
	}

	for i := range channels {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

// buffered channel 使用 demo
func bufferedChannel() {
	ch := make(chan int, 5)  // 缓存 channel
	go worker(0, ch)
	ch <- 'a'
	ch <- 'b'
	ch <- 'c'
	ch <- 'd'
	ch <- 'e'
	time.Sleep(time.Millisecond)
}

// 关闭 channel
func channelClose() {
	ch := make(chan int, 5)  // 缓存 channel
	go worker(0, ch)
	ch <- 'a'
	ch <- 'b'
	ch <- 'c'
	ch <- 'd'
	ch <- 'e'
	close(ch)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()   // channel demo

	fmt.Println("Buffered channel")
	//bufferedChannel()

	fmt.Println("Closed channel")
	//channelClose()

}
