package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 创建 worker 函数
func worker(i int, ch chan int) {

	for v := range ch {
		fmt.Printf("worker %d received %d\n", i, v)
		time.Sleep(time.Second)
	}
}

// 创建 worker 并返回 channel
func createWorker(i int) chan int {
	ch := make(chan int)
	go worker(i, ch)
	return ch
}

// 创建生成数据，定时将数据发送到channel
func generator() chan int {
	c := make(chan int)
	i := 0
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			c <- i
			i++
		}
	}()
	return c
}

// c1, c2 定时生成数据
// 当有数据时 w 会消费数据
// 缺点：w 消费过慢的话，n 中的数据没有缓存，会有数据丢失
func selectDemo1() {
	c1 := generator()
	c2 := generator()
	w := createWorker(0)
	hasValue := false
	n := 0
	for true {
		var activeWorker chan int // activeWorker = nil , select 中 nil 将不会被执行
		if hasValue {             // 通过 hasValue 标志将 w 赋值到 activeWorker 中，赋值后 将执行: case activeWorker <- n:
			activeWorker = w
		}
		select {
		case n = <-c1:
			hasValue = true
			//fmt.Println("Received from c1: ", n)
		case n = <-c2:
			hasValue = true
			//fmt.Println("Receive from c2: ", n)
		case activeWorker <- n:
			hasValue = false
			//default:
			//	fmt.Println("No value received")
		}
	}
}

// c1, c2 定时生成数据
// 当有数据时 w 会消费数据
// 通过 Slice 缓存数据, 防止丢失
func selectDemo2() {
	c1 := generator()
	c2 := generator()
	w := createWorker(0)
	tm := time.After(time.Second * 10)
	tick := time.Tick(time.Millisecond * 500)
	var values []int
	for true {
		var activeWorker chan int // activeWorker = nil , select 中 nil 将不会被执行
		var activeValue int
		if len(values) > 0 { // 如果 values 中有值将 w 赋值到 activeWorker 中，赋值后 将执行: case activeWorker <- values[0]:
			activeWorker = w
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(time.Millisecond * 600): // 500ms内没有其他的case执行，将会执行，否则不执行
			fmt.Println("time out of 600ms")
		case <-tick:
			fmt.Println("len(values)=", len(values))
		case <-tm:
			fmt.Println("Bye bye")
			return
		}
	}
}

func main() {
	//selectDemo1()
	selectDemo2()
}
