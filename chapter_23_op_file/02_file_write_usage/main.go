package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func BufioNewWriterUsage() {
	f, err := os.OpenFile("./learnGo/chapter_23/02_file_write_usage/test.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	write := bufio.NewWriter(f)
	reader := bufio.NewReader(f)
	n := 0
	write.WriteString("\n")
	for true {
		n++
		str, err := reader.ReadString('\n')
		write.WriteString(strconv.Itoa(n) + " " + str)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	f.Seek(0, io.SeekEnd)
	write.Flush()
}

func main() {
	/*
	hello laowang
	hello huixian
	hello jiangjiang
	hello bobo
	hello liuxiang
	*/
	BufioNewWriterUsage()
}
