package main

import (
	"fmt"
	"os"
)

func StdoutUsage() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not Go\n",
		"Errors are values\n",
		"Don't panic\n",
	}

	for _, p := range proverbs {
		// 因为 os.Stdout 也实现了 io.Writer
		n, err := os.Stdout.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}
}

/*
标准输入
Please Input you name: laowang
os.Stdin number:  8
Receive data:  laowang
*/

func StdinUsage() {

	in := make([]byte, 10)
	fmt.Printf("Please Input you name: ")
	n, err := os.Stdin.Read(in)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("os.Stdin number: ", n)
	fmt.Println("Receive data: ", string(in))
}

func main() {
	//StdoutUsage()
	StdinUsage()
}
