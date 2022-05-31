package main

import (
	"context"
	"fmt"
	"time"
)

const ctxKey = "accountId"

func main() {
	c := context.WithValue(context.Background(), ctxKey, "111102020102")
	c, cancel := context.WithTimeout(c, 10 * time.Second)
	go mainTask(c)

	var cmd string

	for true {
		scan, err := fmt.Scan(&cmd)
		if err != nil {
			panic(err)
		}
		fmt.Printf("scan: %d\n", scan)
		if err != nil {
			return
		}
		if cmd == "c" {
			cancel()   // 主任务通过调用 主cancel()，取消子任务执行
		}
	}

	time.Sleep(time.Hour)
}

func mainTask(c context.Context) {
	go func() {
		c1, cancel := context.WithTimeout(c, 10*time.Second)  // 时间超时 和 cancel() 都可以使子任务收到 c.Done()
		defer cancel()
		smallTask(c1, "task1", 9*time.Second)
	}()
	smallTask(c, "task2", 8*time.Second)

	//time.Sleep(4 * time.Second)
	//fmt.Printf("e")
	//cancel()
}

func smallTask(c context.Context, name string, duration time.Duration) {
	fmt.Printf("%s started: %q\n", name, c.Value(ctxKey))
	select {
	case <-time.After(duration):
		fmt.Printf("%s done\n", name)
	case <-c.Done():             // 任务通过接听 c.Done()，判读主任务是否要取消
		fmt.Printf("%s cancelled\n", name)
	}
}
