package main

import (
	"fmt"
	"sync"
	"time"
)

func CondDemo() {
	cond := sync.NewCond(&sync.Mutex{})

	go func() {
		cond.L.Lock()
		fmt.Println("before wait 1")
		cond.Wait()
		fmt.Println("after wait 1")
		cond.L.Unlock()
	}()

	go func() {
		cond.L.Lock()
		fmt.Println("before wait 2")
		cond.Wait()
		fmt.Println("after wait 2")
		cond.L.Unlock()
	}()
	time.Sleep(time.Second * 2)
	cond.Broadcast()
	time.Sleep(time.Second * 1)

}

func CondDemoSignal() {
	cond := sync.NewCond(&sync.Mutex{})

	go func() {
		cond.L.Lock()
		fmt.Println("before wait 1")
		cond.Wait()
		fmt.Println("after wait 1")
		cond.L.Unlock()
	}()

	go func() {
		cond.L.Lock()
		fmt.Println("before wait 2")
		cond.Wait()
		fmt.Println("after wait 2")
		cond.L.Unlock()
	}()

	time.Sleep(time.Second * 1)
	cond.Signal()
	time.Sleep(time.Second * 1)
	cond.Signal()
	time.Sleep(time.Second * 1)

}



func main() {
	//CondDemo()
	CondDemoSignal()
}
