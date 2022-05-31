package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

/*IOCopyUsage
io.Copy() 可以轻松地将数据从一个 Reader 拷贝到另一个 Writer。
它抽象出 for 循环模式并正确处理 io.EOF 和 字节计数
*/
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


// CopyFileToStdout 使用 io.Copy() 函数重写从文件读取并打印到标准输出
func CopyFileToStdout() {
	file, err := os.Open("./proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	if _, err := io.Copy(os.Stdout, file); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


/*WriteStringToFile
io.WriteString()，此函数让我们方便地将字符串类型写入一个 Writer：
*/
func WriteStringToFile() {
	file, err := os.Create("./magic_msg.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	if _, err := io.WriteString(file, "Go is fun!"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}



func main() {
	//IOCopyUsage()
	//CopyFileToStdout()
	WriteStringToFile()
}
