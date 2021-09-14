package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicRWMutexInt struct {
	value int
	lock  *sync.RWMutex
}

func (a *atomicRWMutexInt) increment() {

	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		time.Sleep(time.Second)
		fmt.Println("Execute increment() function")
		a.value++
	}()
}

func (a *atomicRWMutexInt) get() int {
	a.lock.RLock()
	defer a.lock.RUnlock()
	time.Sleep(time.Second * 1)
	fmt.Printf("Execute get() function: %d\n", a.value)
	return a.value
}

func RWMutexDemo() {
	a := &atomicRWMutexInt{
		value: 0,
		lock:  &sync.RWMutex{},
	}
	go func() {
		for i := 0; i < 2; i++ {
			a.increment()
		}
	}()
	time.Sleep(time.Millisecond*10)
	for i := 0; i < 3; i++ {
		go a.get()
	}
	time.Sleep(time.Second * 7)
}

func main() {
	RWMutexDemo()
}
