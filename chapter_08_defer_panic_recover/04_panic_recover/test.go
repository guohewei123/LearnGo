package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Printf("Error occured: %s\n", err.Error())
		} else {
			panic(fmt.Sprintf("I don't know what to do: %v", r))
		}
	}()

	// 测试一 输出: Error occured: this is a error
	//panic(errors.New("this is a error"))

	// 测试二 输出: Error occured: runtime error: integer divide by zero
	/*a := 0
	b := 1/a
	fmt.Println(b)*/

	/* 测试 三 输出:
	panic: 123 [recovered]
	        panic: I don't know what to do: 123
	*/
	panic(123)
}

func main() {
	tryRecover()
}
