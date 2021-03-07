package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

// for 省略初始条件,相当于while
func convertTOBin(n int) string {
	inputN := n
	res := ""
	for ; n >0; n /= 2 {
		lsb := n % 2
		res = strconv.Itoa(lsb) + res
	}
	fmt.Printf("Int: %10d, Bin: %s\n", inputN, res)
	return res
}

// for 省略初始条件和递增条件
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for ; scanner.Scan(); {
		fmt.Println(scanner.Text())
	}
}

//3.初始条件,结束条件,递增表达式都不加就是死循环
func forever() string {
	for {
		fmt.Println("Forever loop")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	convertTOBin(5)
	convertTOBin(9)
	convertTOBin(13)
	convertTOBin(32)
	convertTOBin(34534)
	fmt.Println()
	const filename = "learnGo/chapter_02/02_branch/abc.txt" // abc.txt = aaaa\nbbbbb\ncccc
	printFile(filename)
	fmt.Println()
	forever()
}
