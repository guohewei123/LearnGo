package main

import (
	"fmt"
	"io"
	"os"
)

func OsFileWriteUsage() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not Go\n",
		"Errors are values\n",
		"Don't panic\n",
	}
	file, err := os.Create("./proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for _, p := range proverbs {
		// file 类型实现了 io.Writer
		n, err := file.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}
	fmt.Println("file write done")
}

/*
同时，io.File 也可以用作读取器来从本地文件系统读取文件的内容。
例如，下面的例子展示了如何读取文件并打印其内容：
*/

func OsFileReadUsage() {
	file, err := os.Open("./proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	p := make([]byte, 4)
	for {
		n, err := file.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
}

func main() {
	//OsFileWriteUsage()
	OsFileReadUsage()
}
