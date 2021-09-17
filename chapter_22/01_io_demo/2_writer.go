package main

import (
	"bytes"
	"fmt"
	"os"
)

/* -------------------------------
1. 它使用 bytes.Buffer 类型作为 io.Writer 将数据写入内存缓冲区。
输出：
Channels orchestrate mutexes serializeCgo is not GoErrors are valuesDon't panic
------------------------------- */

func WriterExample() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize",
		"Cgo is not Go",
		"Errors are values",
		"Don't panic",
	}
	var writer bytes.Buffer

	for _, p := range proverbs {
		// Write() 方法有两个返回值，一个是写入到目标资源的字节数，一个是发生错误时的错误。
		n, err := writer.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}

	fmt.Println(writer.String())
}

/* -------------------------------
2. 自己实现一个 Writer
   下面我们来实现一个名为 chanWriter 的自定义 io.Writer ，它将其内容作为字节序列写入 channel 。
输出：
Stream me!
------------------------------- */
type chanWriter struct {
	// ch 实际上就是目标资源
	ch chan byte
}

func newChanWriter() *chanWriter {
	return &chanWriter{make(chan byte, 1024)}
}

func (w *chanWriter) Chan() <-chan byte {
	return w.ch
}

func (w *chanWriter) Write(p []byte) (int, error) {
	n := 0
	// 遍历输入数据，按字节写入目标资源
	for _, b := range p {
		w.ch <- b
		n++
	}
	return n, nil
}

func (w *chanWriter) Close() error {
	close(w.ch)
	return nil
}


/* -------------------------------
要使用这个 Writer，只需在函数 main() 中调用 writer.Write()（在单独的goroutine中）。
因为 chanWriter 还实现了接口 io.Closer ，所以调用方法 writer.Close() 来正确地关闭channel，以避免发生泄漏和死锁。
输出：
Stream me!
------------------------------- */

func SelfWriter() {
	writer := newChanWriter()
	go func() {
		defer writer.Close()
		writer.Write([]byte("Stream "))
		writer.Write([]byte("me!"))
	}()
	for c := range writer.Chan() {
		fmt.Printf("%c", c)
		//fmt.Printf("%c\n", c)
	}
	fmt.Println()
}

func main() {
	//WriterExample()
	SelfWriter()
}
