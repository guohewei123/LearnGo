package main

import "fmt"

// 通过指针来交换值
func swap1(a, b *int) {
	*a, *b = *b, *a
}

func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 3, 4
	swap1(&a, &b)
	fmt.Println(a, b)
	a, b = 3, 4
	a, b = swap2(a, b)
	fmt.Println(a, b)
}
