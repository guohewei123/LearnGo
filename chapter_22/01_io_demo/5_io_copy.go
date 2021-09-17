package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func IOCopyUsage() {
	proverbs := new(bytes.Buffer)
	proverbs.WriteString("Channels orchestrate mutexes serialize\n")
	proverbs.WriteString("Cgo is not Go\n")
	proverbs.WriteString("Errors are values\n")
	proverbs.WriteString("Don't panic\n")

	file, err := os.Create("./proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// io.Copy 完成了从 proverbs 读取数据并写入 file 的流程
	if _, err := io.Copy(file, proverbs); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("file created")
}

func main() {
	IOCopyUsage()
}
