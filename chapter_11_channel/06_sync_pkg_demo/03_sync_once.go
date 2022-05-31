package main

import (
	"fmt"
	"sync"
)

func PrintHelloWorld() {
	fmt.Println("Execute PrintHelloWorld() function")
}

func OnceDemo() {
	a := &sync.Once{}
	for i := 1; i < 10; i++ {
		a.Do(PrintHelloWorld)
		//PrintHelloWorld()
	}
}

func main() {
	OnceDemo()
}
