package main

import (
	"fmt"
	"queue/queue"
)

// 方法二： 通过别名扩展系统包
func main() {
	q := queue.Queue{1}
	q.Append(2)
	q.Append(3)
	fmt.Println(q)
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
