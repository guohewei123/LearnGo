## 22-1 io包的使用
原文：https://medium.com/learning-the-go-programming-language/streaming-io-in-go-d93507931185

译文：Go 中 io 包的使用方法：https://segmentfault.com/a/1190000015591319

### 前言
在 Go 中，输入和输出操作是使用原语实现的，这些原语将数据模拟成可读的或可写的字节流。
为此，Go 的 io 包提供了 io.Reader 和 io.Writer 接口，分别用于数据的输入和输出，如图：

![](.README_images/6b493e48.png)

Go 官方提供了一些 API，支持对内存结构，文件，网络连接等资源进行操作
本文重点介绍如何实现标准库中 io.Reader 和 io.Writer 两个接口，来完成流式传输数据。


### 1. io.Reader

io.Reader 表示一个读取器，它将数据从某个资源读取到传输缓冲区。在缓冲区中，数据可以被流式传输和使用。
如图：

![](.README_images/2e198c18.png)

对于要用作读取器的类型，它必须实现 io.Reader 接口的唯一一个方法 Read(p []byte)。
换句话说，只要实现了 Read(p []byte) ，那它就是一个读取器。

  ```go
  type Reader interface{ 
    Read(p []byte) (n int, err error)
  }   
  ```
Read() 方法有两个返回值，一个是读取到的字节数，一个是发生错误时的错误。
同时，如果资源内容已全部读取完毕，应该返回 io.EOF 错误。

利用 Reader 可以很容易地进行流式数据传输。Reader 方法内部是被循环调用的，每次迭代，它会从数据源读取一块数据放入缓冲区 p （即 Read 的参数 p）中，直到返回 io.EOF 错误时停止。

### 2. io.Writer
io.Writer 表示一个编写器，它从缓冲区读取数据，并将数据写入目标资源。

![](.README_images/b13c44d3.png)

对于要用作编写器的类型，必须实现 io.Writer 接口的唯一一个方法 Write(p []byte)
同样，只要实现了 Write(p []byte) ，那它就是一个编写器。

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```
Write() 方法有两个返回值，一个是写入到目标资源的字节数，一个是发生错误时的错误。
标准库提供了许多已经实现了 io.Writer 的类型。

- 可以使用 bytes.Buffer 类型作为 io.Writer 将数据写入内存缓冲区
  ```go
  func main() {
      proverbs := []string{
          "Channels orchestrate mutexes serialize",
          "Cgo is not Go",
          "Errors are values",
          "Don't panic",
      }
      var writer bytes.Buffer
  
      for _, p := range proverbs {
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
  ```

  ```
  输出打印的内容：
  Channels orchestrate mutexes serializeCgo is not GoErrors are valuesDon't panic
  ```
  
- 自己实现一个 Writer
  
  下面我们来实现一个名为 chanWriter 的自定义 io.Writer ，它将其内容作为字节序列写入 channel 。

  ```go
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
  
  func main() {
      writer := newChanWriter()
      go func() {
          defer writer.Close()
          writer.Write([]byte("Stream "))
          writer.Write([]byte("me!"))
      }()
      for c := range writer.Chan() {
          fmt.Printf("%c", c)
      }
      fmt.Println()
  }
  ```
  要使用这个 Writer，只需在函数 main() 中调用 writer.Write()（在单独的goroutine中）。
  因为 chanWriter 还实现了接口 io.Closer ，所以调用方法 writer.Close() 来正确地关闭channel，以避免发生泄漏和死锁。
### 3. io 包里其他有用的类型和方法

#### 3.1 os.File
类型 os.File 表示本地系统上的文件。它实现了 io.Reader 和 io.Writer ，因此可以在任何 io 上下文中使用。
例如，下面的例子展示如何将连续的字符串切片直接写入文件：
```go
func main() {
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
```
同时，io.File 也可以用作读取器来从本地文件系统读取文件的内容。
例如，下面的例子展示了如何读取文件并打印其内容：

```go
func main() {
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
```

#### 3.2 标准输入、输出和错误
os 包有三个可用变量 os.Stdout ，os.Stdin 和 os.Stderr ，它们的类型为 *os.File，分别代表 系统标准输入，系统标准输出 和 系统标准错误 的文件句柄。
例如，下面的代码直接打印到标准输出：
```go
func main() {
    proverbs := []string{
        "Channels orchestrate mutexes serialize\n",
        "Cgo is not Go\n",
        "Errors are values\n",
        "Don't panic\n",
    }

    for _, p := range proverbs {
        // 因为 os.Stdout 也实现了 io.Writer
        n, err := os.Stdout.Write([]byte(p))
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        if n != len(p) {
            fmt.Println("failed to write data")
            os.Exit(1)
        }
    }
}
```

#### 3.3 io.Copy()

io.Copy() 可以轻松地将数据从一个 Reader 拷贝到另一个 Writer。
它抽象出 for 循环模式（我们上面已经实现了）并正确处理 io.EOF 和 字节计数。
下面是我们之前实现的简化版本：

```go
func main() {
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
```

那么，我们也可以使用 io.Copy() 函数重写从文件读取并打印到标准输出的先前程序，如下所示：

```go
func main() {
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
```

#### 3.4 io.WriteString()
此函数让我们方便地将字符串类型写入一个 Writer：
```go
func main() {
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
```

#### 3.5 使用管道的 Writer 和 Reader

类型 io.PipeWriter 和 io.PipeReader 在内存管道中模拟 io 操作。
数据被写入管道的一端，并使用单独的 goroutine 在管道的另一端读取。
下面使用 io.Pipe() 创建管道的 reader 和 writer，然后将数据从 proverbs 缓冲区复制到io.Stdout ：

```go
func main() {
    proverbs := new(bytes.Buffer)
    proverbs.WriteString("Channels orchestrate mutexes serialize\n")
    proverbs.WriteString("Cgo is not Go\n")
    proverbs.WriteString("Errors are values\n")
    proverbs.WriteString("Don't panic\n")

    piper, pipew := io.Pipe()

    // 将 proverbs 写入 pipew 这一端
    go func() {
        defer pipew.Close()
        io.Copy(pipew, proverbs)
    }()

    // 从另一端 piper 中读取数据并拷贝到标准输出
    io.Copy(os.Stdout, piper)
    piper.Close()
}
```

#### 3.6 缓冲区 io

标准库中 bufio 包支持 缓冲区 io 操作，可以轻松处理文本内容。
例如，以下程序逐行读取文件的内容，并以值 '\n' 分隔：

```go
func main() {
    file, err := os.Open("./planets.txt")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()
    reader := bufio.NewReader(file)

    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            if err == io.EOF {
                break
            } else {
                fmt.Println(err)
                os.Exit(1)
            }
        }
        fmt.Print(line)
    }

}
```

#### 3.7 ioutil
io 包下面的一个子包 utilio 封装了一些非常方便的功能
例如，下面使用函数 ReadFile 将文件内容加载到 []byte 中。

```go

package main

import (
  "fmt"
  "io/ioutil"
  "os"
)

func main() {
    bytes, err := ioutil.ReadFile("./planets.txt")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Printf("%s", bytes)
}
```

### Seeker
  ```go
  type Seeker interface{ 
    Seek(offset int64, whence int) (ret int64, err error)
  }   
  ```
  - Seek 设置下一次读写操作的指针位置，每次的读写都是从指针位置开始的
  - whence 为0：表示从数据的开头开始移动指针
  - whence 为1：表示从数据当前位置开始移动指针
  - whence 为2：表示从数据的尾部开始移动指针
  - offset 是指针移动的偏移量
    
### Closer
    ```go
    type Closer interface{ 
      Close() error
    }   
    ```
    - Close 用于关闭文件，关闭连接，关闭数据库等
    
