package main

import "fmt"

func main() {
	//fmtPrintString()
	fmtPrintStrut()
}

/* fmtPrintString
1. "%s" "%q"拼接对比
输出结果：
print: 测试数据
print: "测试数据"
*/
func fmtPrintString() {
	a := "测试数据"
	fmt.Printf("print: %s\n", a)
	fmt.Printf("print: %q\n", a)
}

/* fmtPrintStrut
2. "%v" "%+v" 拼接对比
输出结果：
print: {张三 18}
print: {name:张三 age:18}
*/
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
