package main

import "fmt"

// 方法一: 使用闭包实现 adder
func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

// 方法二 "正统" 函数式编程 实现 adder
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	fmt.Println("---- 使用闭包实现 adder --- ")
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d \n", i, a(i))
	}

	fmt.Println(`---- "正统" 函数式编程 实现 adder ----`)
	a2 := adder2(0)
	for i := 0; i < 10; i++ {
		var sum int
		sum, a2 = a2(i)
		fmt.Printf("0 + 1 + ... + %d = %d \n", i, sum)
	}
}

// 输出结果
/*
---- 使用闭包实现 adder ---
0 + 1 + ... + 0 = 0
0 + 1 + ... + 1 = 1
0 + 1 + ... + 2 = 3
0 + 1 + ... + 3 = 6
0 + 1 + ... + 4 = 10
0 + 1 + ... + 5 = 15
0 + 1 + ... + 6 = 21
0 + 1 + ... + 7 = 28
0 + 1 + ... + 8 = 36
0 + 1 + ... + 9 = 45
---- "正统" 函数式编程 实现 adder ----
0 + 1 + ... + 0 = 0
0 + 1 + ... + 1 = 1
0 + 1 + ... + 2 = 3
0 + 1 + ... + 3 = 6
0 + 1 + ... + 4 = 10
0 + 1 + ... + 5 = 15
0 + 1 + ... + 6 = 21
0 + 1 + ... + 7 = 28
0 + 1 + ... + 8 = 36
0 + 1 + ... + 9 = 45
*/
