package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 函数定义，可以返回多个值
func eval(a, b int, op string) (int, error) {
	res := 0
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		//res = a / b
		res, _ = div(a, b)
	default:
		//return res, fmt.Errorf("unsupport operation: %s", op)
	}
	return res, nil
}

// 7 / 4 = 1 ... 3
// 输出起名字
func div(a, b int) (q, r int) {
	//q = a / b
	//r = a % b
	//fmt.Printf("%d / %d = %d ... %d\n", a, b, q, r)
	//return
	return a / b, a % b
}

// 函数的参数可以是函数
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function: %s with args: (%d, %d)\n", opName, a, b)
	return op(a, b)
}

// 定义pow函数
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表
func sum(numbers ...int) int {
	sum := 0
	for i:=range numbers {
		sum += numbers[i]
	}
	return sum
}

func main() {
	if res, err := eval(4, 3, "-"); err != nil {
		fmt.Printf("Error: %s\n", err)
		panic(err)
	} else {
		fmt.Println(res)
	}

	q, r := div(10, 3)
	fmt.Printf("10 / 3 = %d ... %d\n", q, r)

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))
}
