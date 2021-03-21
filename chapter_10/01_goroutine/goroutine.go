package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var sum [20]int
	for i := 0; i < 20; i++ {
		go func(ii int) {
			//fmt.Printf("Hello from goroutine %d\n", ii)
			for {
				sum[ii] ++
				runtime.Gosched() // 手动交出控制权
			}
		}(i)
	}
	time.Sleep(time.Second * 20)
	fmt.Println(sum)
}
