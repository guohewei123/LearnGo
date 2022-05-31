package queue

type Queue []int

func (q *Queue) Append(val int) {
	*q = append(*q, val)
}

func (q *Queue) Pop() int {
	top := (*q)[0]
	*q = (*q)[1:]
	return top
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
