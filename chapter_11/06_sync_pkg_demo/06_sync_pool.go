package main

import (
	"fmt"
	"sync"
)

func PoolDemo() {
	pool := &sync.Pool{}
	for i := 0; i < 10; i++ {
		pool.Put(i)
	}

	for i := 0; i < 10; i++ {
		val := pool.Get()
		fmt.Println(val)
		//pool.Put(val)
	}


}


func main() {
	PoolDemo()
}
