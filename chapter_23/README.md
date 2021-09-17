## goland 文件操作

### 1. os 包常量

// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
O_RDONLY int = syscall.O_RDONLY // open the file read-only.
O_WRONLY int = syscall.O_WRONLY // open the file write-only.
O_RDWR   int = syscall.O_RDWR   // open the file read-write.
// The remaining values may be or'ed in to control behavior.
O_APPEND int = syscall.O_APPEND // append data to the file when writing.
O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.     和 O_CREATE 配合使用，文件必须不存在
O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.                    打开文件用于同步IO
O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.  如果可能打开时清空文件


### 2. 读取文件方法

- os.openFile() 用户打开文件
  ```go
  func Create(name string) (*File, error) {
      return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
  }
  ```
- bufio.NewReader(f)  将文件转化为Reader
  ```go
  // NewReader returns a new Reader whose buffer has the default size.
  func NewReader(rd io.Reader) *Reader {
  return NewReaderSize(rd, defaultBufSize)
  }
  ```
- reader.ReadString("string") 调用reader上的方法 还有 ReadLine ReadByte ReadSlice
- ioutil.ReadAll(f) 直接读取整个文件 os.ReadFile(文件路径) 也能达到同样效果
- os.ReadDir("./) 读取文件夹，获取目标文件夹下的文件信息

### 3. 写文件的方法
- os.OpenFile() 用于打开文件获取 *file
- f.Seek()  移动光标位置
- f.WriteString() 直接写入
- bufio.NewWrite(f) 创建一个缓存的写
  - writer.WriteString() 写入内存
  - write.Flush() 缓存内容生效，写入文件

### 4. 复制文件