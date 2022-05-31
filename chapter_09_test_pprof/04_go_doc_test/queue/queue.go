package myQueue

// FIFO 先入先出 队列
// An FIFO queue
type Queue []int

// 添加元素到队列中 pushes the element into the queue.
// 		例如: q.Push(123)
func (q *Queue) Append(val int) {
	*q = append(*q, val)
}

// 弹出队列中元素
// pop the element of the queue.
func (q *Queue) Pop() int {
	top := (*q)[0]
	*q = (*q)[1:]
	return top
}

// 判断队列是否为空
// Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
