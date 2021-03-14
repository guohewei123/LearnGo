package main

import (
	anyQueue "anyQueue/any_queue"
	"fmt"
)

// 方法二： 通过别名扩展系统包
func main() {
	q := anyQueue.Queue{1}
	q.Append(2)
	q.Append(3)
	fmt.Println(q)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	q.Append("你好")
	fmt.Println(q.Pop())
	q.Append("Hello world")
	fmt.Println(q.Pop())
}
