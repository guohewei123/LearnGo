package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// 1 1 2 3 5 8 13 ...
func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 定义斐波那契返回函数，用于实现 Reader 接口
type intGen func() int

// 为 intGen 实现 Reader 接口
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	s := strconv.Itoa(next) + "\n"
	//s := fmt.Sprintf("%d\n", next)
	if next > 100 {
		return 0, io.EOF
	}
	return strings.NewReader(s).Read(p)
	//s := fmt.Sprint("%s")

}

// 读取实现 Reader 接口的对象
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {

	f := fibonacci()
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())

	fmt.Println("---------- 使用 printFileContents 打印 ------------")
	printFileContents(f)

}

// 输出结果
/*
---------- 使用 printFileContents 打印 ------------
1
1
2
3
5
8
13
21
34
55
89
*/
