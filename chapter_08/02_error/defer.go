package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

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
	filename := "learnGo/chapter_08/02_error/abc.txt"
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		/* 通过实现该接口，也可以自定义error
		type error interface {
			Error() string
		}
		*/
		err = errors.New("this is a custom err!!!")  // 使用 errors.New 自定义 error
		if pathError, ok := err.(*os.PathError); !ok {  // 判断 err 是否是 PathError
			fmt.Println("Error: ", err.Error())        // Error:  this is a custom err!!!
			panic(err)
		}else {
			// 输出:    'Operation: open, Path: learnGo/chapter_08/02_error/abc.txt, Error: file exists'
			fmt.Printf("Operation: %s, Path: %s, Error: %s\n", pathError.Op, pathError.Path, pathError.Err)
			return
		}
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
	fmt.Println()
	writerFile()
}
