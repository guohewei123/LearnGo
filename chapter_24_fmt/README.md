# Golang 字符串拼接

1. "%s" "%q" 拼接对比
```go
func fmtPrintString() {
    a := "测试数据"
    fmt.Printf("print: %s\n", a)
    fmt.Printf("print: %q\n", a)
}

/*
输出结果：
print: 测试数据
print: "测试数据"
*/
```
2. "%v" "%+v" 拼接对比
```go
func fmtPrintStrut() {
	var a = struct {
		name string
		age  int
	}{
		name: "张三",
		age:  18,
	}
	fmt.Printf("print: %v\n", a)
	fmt.Printf("print: %+v\n", a)
}

/* 
输出结果：
print: {张三 18}
print: {name:张三 age:18}
*/
```