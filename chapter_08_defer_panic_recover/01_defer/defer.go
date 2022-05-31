package main

import (
	"bufio"
	"fmt"
	"os"
)

// 测试 defer 的使用方法、调用顺序、参数在 defer 语句时计算
func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
	defer fmt.Println(4)
}

/* 执行 tryDefer() 输出结果
3
2
1
*/

// 定义一个斐波那契数列生成函数
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 使用 defer 写文件
func writerFile() {
	file, err := os.Create("learnGo/chapter_08/01_defer/abc.txt")
	if err != nil {
		fmt.Println("Create file failed!")
		panic(err)
	}
	defer file.Close()   // 创建好文件就 defer Close() 确保函数退出时关闭文件
	writer := bufio.NewWriter(file)  // 使用缓存写文件 先写到内存
	defer writer.Flush()             // 当函数退出时 将内存中的文件写到真实的文件中
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())   // 将 fib 数列写入到 buffer file 中
	}
}

func main() {
	tryDefer()
	writerFile()
}
