package main

import (
	"context"
	"fmt"
	"time"
)

/* -------------------------------
1. 使用chan实现异步接收
------------------------------- */
func chanReceive(bufChan chan int, flag chan bool) {
	t := time.Tick(time.Millisecond * 500)
	for _ = range t {
		select {
		case m := <-bufChan:
			fmt.Println("接收到了: ", m)
		case <-flag:
			fmt.Println("结束接收")
			return
		}
	}
}

func chanSendReceive() {
	bufChan := make(chan int)
	flag := make(chan bool)
	go chanReceive(bufChan, flag)
	for i := 0; i < 10; i++ {
		bufChan <- i
	}
	fmt.Println("结束发送")
	flag <- true
	time.Sleep(time.Second)
	fmt.Println("game over")
}

/* -------------------------------
2. 使用WithCancel实现异步接收
------------------------------- */
func ctxReceive(bufChan chan int, ctx context.Context) {
	t := time.Tick(time.Millisecond * 500)
	for _ = range t {
		select {
		case m := <-bufChan:
			fmt.Println("接收到了: ", m)
		case <-ctx.Done():
			fmt.Println("结束接收")
			return
		}
	}
}

func ctxWithCancel() {
	bufChan := make(chan int)
	ctx, clear := context.WithCancel(context.Background())
	go ctxReceive(bufChan, ctx)
	for i := 0; i < 10; i++ {
		bufChan <- i
	}
	fmt.Println("结束发送")
	clear()
	time.Sleep(time.Second)
	fmt.Println("game over")
}

/* -------------------------------
3. 使用WithCancel + WithValue 实现异步接收
------------------------------- */
func ctxWithCancelAndWithValue() {
	bufChan := make(chan int)
	ctx := context.WithValue(context.Background(), "testKey", "testVal")
	ctx, clear := context.WithCancel(ctx)
	go ctxWithValueReceive(bufChan, ctx)
	for i := 0; i < 10; i++ {
		bufChan <- i
	}
	fmt.Println("结束发送")
	clear()
	time.Sleep(time.Second)
	fmt.Println("game over")
}

func ctxWithValueReceive(bufChan chan int, ctx context.Context) {
	t := time.Tick(time.Millisecond * 500)
	for _ = range t {
		select {
		case m := <-bufChan:
			fmt.Printf("接收到了: %d, ctx.Value(testKey)=%s\n", m, ctx.Value("testKey"))
		case <-ctx.Done():
			fmt.Println("结束接收, ctx.Value(testKey): ", ctx.Value("testKey"))
			return
		}
	}
}

/* -------------------------------
4. 使用 WithDeadline + WithValue  实现异步接收
------------------------------- */
func ctxWithDeadlineAndWithValue() {
	bufChan := make(chan int)
	ctx := context.WithValue(context.Background(), "testKey", "testVal")
	ctx, cancelFunc := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	go ctxWithValueReceive(bufChan, ctx)
	for i := 0; i < 10; i++ {
		//if i == 2{
		//	cancelFunc() // 强制停止, 不等待Deadline
		//}
		bufChan <- i
	}
	fmt.Println("结束发送")
	defer cancelFunc()   // 函数退出时强制停止ctx, 不等待Deadline
	time.Sleep(time.Second)
	fmt.Println("game over")
}

/* -------------------------------
5. 使用 WithTimeout + WithValue  实现异步接收
------------------------------- */
func ctxWithTimeoutAndWithValue() {
	bufChan := make(chan int)
	ctx := context.WithValue(context.Background(), "testKey", "testVal")
	ctx, cancelFunc := context.WithTimeout(ctx, time.Second*2)
	go ctxWithValueReceive(bufChan, ctx)
	for i := 0; i < 10; i++ {
		//if i == 2{
		//	cancelFunc() // 强制停止, 不等待Deadline
		//}
		bufChan <- i
	}
	fmt.Println("结束发送")
	defer cancelFunc()   // 函数退出时强制停止ctx, 不等待Deadline
	time.Sleep(time.Second)
	fmt.Println("game over")
}


func main() {
	//chanSendReceive()
	//ctxWithCancel()
	//ctxWithCancelAndWithValue()
	//ctxWithDeadlineAndWithValue()
	ctxWithTimeoutAndWithValue()
}
