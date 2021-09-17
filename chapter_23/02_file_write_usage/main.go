package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func BufioNewWriterUsage() {
	f, err := os.OpenFile("./learnGo/chapter_23/02_file_write_usage/test1.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	wirter := bufio.NewWriter(f)
	reader := bufio.NewReader(f)
	n := 0
	for true {
		n++
		str, err := reader.ReadString('\n')
		wirter.WriteString(strconv.Itoa(n) + " " + str)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	f.Seek(0, 0)
	wirter.Flush()

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
