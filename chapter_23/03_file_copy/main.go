package main

import (
	"io"
	"os"
)

func IoCopyUsage() {
	f1, _ := os.OpenFile("./learnGo/chapter_23/03_file_copy/test1.txt", os.O_CREATE|os.O_RDWR, 0777)
	f2, _ := os.OpenFile("./learnGo/chapter_23/03_file_copy/test2.txt", os.O_CREATE|os.O_RDWR, 0777)
	io.Copy(f2, f1)

}

func main() {
	/*
	hello laowang
	hello huixian
	hello jiangjiang
	hello bobo
	hello liuxiang
	*/
	IoCopyUsage()
}
