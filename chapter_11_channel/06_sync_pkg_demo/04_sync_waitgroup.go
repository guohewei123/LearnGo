package main

import (
	"fmt"
	"sync"
	"time"
)

func blockPrint(n int, w *sync.WaitGroup) {
	go func() {
		time.Sleep(time.Second * time.Duration(n))
		fmt.Printf("Execute block() function: %d\n", n)
		w.Done()
	}()
}

func WaitGroupDemo() {
	w := &sync.WaitGroup{}
	w.Add(5)
	for i := 0; i < 5; i++ {
		blockPrint(i, w)
	}
	w.Wait()
	fmt.Println("结束")
}

func main() {
	WaitGroupDemo()
}
