package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  *sync.Mutex
}

func (a *atomicInt) increment() {

	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		time.Sleep(time.Second)
		fmt.Println("Execute increment() function")
		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	fmt.Println("Execute get() function")
	return a.value
}

func mutexDemo() {
	a := &atomicInt{
		value: 0,
		lock:  &sync.Mutex{},
	}
	a.increment()
	go func() {
		a.increment()
	}()
	fmt.Println(a.get())
}

func main() {
	mutexDemo()
}
