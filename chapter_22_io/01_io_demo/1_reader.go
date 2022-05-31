package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/* -------------------------------
1. 通过 string.NewReader(string) 创建一个字符串读取器, 然后流式地按字节读取
输出：
6 Clear
6 is bet
6 ter th
6 an cle
3 ver
EOF:  0
------------------------------- */
func stringReaderDemo() {
	reader := strings.NewReader("Clear is better than clever")
	p := make([]byte, 6)

	for true {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF: ", n)
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(n, string(p))
		p = make([]byte, 6) // 清空p, 防止最后一次不足6个字符，返回上一次的缓存值
	}
}

/* -------------------------------
2. 自己实现一个 Reader
利用 Reader 可以很容易地进行流式数据传输。Reader 方法内部是被循环调用的，每次迭代，
它会从数据源读取一块数据放入缓冲区 p （即 Read 的参数 p）中，直到返回 io.EOF 错误时停止。
输出：
HelloItsamwhereisthesun
------------------------------- */
type selfAlphaReader struct {
	// 资源
	src string
	// 当前读取到的位置
	cur int
}

// 创建一个实例
func newSelfAlphaReader(src string) *selfAlphaReader {
	return &selfAlphaReader{src: src}
}

// 过滤函数
func selfAlpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

// Read 方法
func (a *selfAlphaReader) Read(p []byte) (int, error) {
	// 当前位置 >= 字符串长度 说明已经读取到结尾 返回 EOF
	if a.cur >= len(a.src) {
		return 0, io.EOF
	}

	// x 是剩余未读取的长度
	x := len(a.src) - a.cur
	n, bound := 0, 0
	if x >= len(p) {
		// 剩余长度超过缓冲区大小，说明本次可完全填满缓冲区
		bound = len(p)
	} else if x < len(p) {
		// 剩余长度小于缓冲区大小，使用剩余长度输出，缓冲区不补满
		bound = x
	}

	buf := make([]byte, bound)
	for n < bound {
		// 每次读取一个字节，执行过滤函数
		if char := selfAlpha(a.src[a.cur]); char != 0 {
			buf[n] = char
		}
		n++
		a.cur++
	}
	// 将处理后得到的 buf 内容复制到 p 中
	copy(p, buf)
	return n, nil
}

func testSelfAlphaReader() {
	reader := newSelfAlphaReader("Hello! It's 9am, where is the sun?")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()
}

/* -------------------------------
3.标准库已经实现了许多 Reader。
使用一个 Reader 作为另一个 Reader 的实现是一种常见的用法。
这样做可以让一个 Reader 重用另一个 Reader 的逻辑，下面展示通过更新 alphaReader 以接受 io.Reader 作为其来源。
输出：
HelloItsamwhereisthesun
------------------------------- */
type alphaReader struct {
	// alphaReader 里组合了标准库的 io.Reader
	reader io.Reader
}

func newAlphaReader(reader io.Reader) *alphaReader {
	return &alphaReader{reader: reader}
}

func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func (a *alphaReader) Read(p []byte) (int, error) {
	// 这行代码调用的就是 io.Reader
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := alpha(p[i]); char != 0 {
			buf[i] = char
		}
	}

	copy(p, buf)
	return n, nil
}

func testAlphaReader() {
	//  使用实现了标准库 io.Reader 接口的 strings.Reader 作为实现
	reader := newAlphaReader(strings.NewReader("Hello! It's 9am, where is the sun?"))
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()
}

/*
这样做的另一个优点是 alphaReader 能够从任何 Reader 实现中读取。
例如，以下代码展示了 alphaReader 如何与 os.File 结合以过滤掉文件中的非字母字符：
*/

func alphaReaderFile() {
	// file 也实现了 io.Reader
	file, err := os.Open("./learnGo/chapter_22/01_io_demo/1_reader.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// 任何实现了 io.Reader 的类型都可以传入 newAlphaReader
	// 至于具体如何读取文件，那是标准库已经实现了的，我们不用再做一遍，达到了重用的目的
	reader := newAlphaReader(file)
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()
}


func main() {
	//stringReaderDemo()
	//testSelfAlphaReader()
	//testAlphaReader()
	alphaReaderFile()
}
