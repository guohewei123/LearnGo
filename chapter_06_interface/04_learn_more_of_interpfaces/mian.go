package main

import "fmt"

type GrpcError struct{}

func (e GrpcError) Error() string {
	return "GrpcError"
}

func main() {
	err := cal()
	if err != nil {
		fmt.Println("----", err)
	}
	fmt.Println(err)            // 打印：<nil>
	fmt.Println(err == nil)     // 打印：false
}

func cal() error {
	var err *GrpcError = nil
	return err
}