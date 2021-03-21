package myQueue

import "fmt"

func ExampleQueue_Append() {
	q := Queue{1}
	q.Append(2)
	q.Append(3)
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	// Output:
	// false
	// 1
	// 2
	// false
	// 3
	// true
}