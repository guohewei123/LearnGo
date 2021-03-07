## 一、变量定义
### 1. 使用 var 关键字
- `var a, b, c bool`
- `var s1, s2 string= "hello", "world"`
- 可放在函数内，或直接放在包内
- 使用var()集中定义变量
### 2. 让编译器自动决定
- `var a, b, c = 123, "test", true`
### 3. 使用 := 定义变量
- `a, b, i, s1, s2 := true, false, 2, "hello", "world"`
- 注意: 只能在函数内使用


```go
package main

import (
	"fmt"
)

//函数外面定义变量时,必须使用var关键字,不能使用:=
//这些变量作用域,是包内变量,不存在全局变量说法
var (
	aa = "test"
	bb = true
	cc = 123
)

func variableZeroValue() {
	var a int
	var b string
	fmt.Printf("%d, %q\n", a, b)
}

//定义变量类型,不能写在一行
func variableInitValue() {
	var a, b int = 2, 3
	var c, d string = "hello", "world"
	fmt.Println(a, b, c, d)
}

//省略变量类型,可以写在一行
func varTypeDefvalue() {
	var a, b, c, d = 1, 3, true, "world"
	fmt.Println(a, b, c, d)
}

//省略var,使用 := 来定义
func variableValueShort() {
	a, b, c, d := 1, 3, true, "world"
	b = 5
	fmt.Println(a, b, c, d)
}

func main() {
	variableZeroValue()
	variableInitValue()
	varTypeDefvalue()
	variableValueShort()
	fmt.Println(aa, bb, cc)
}

```

执行输出结果：
```
0, ""
2 3 hello world
1 3 true world
1 5 true world
test true 123
```

## 一、内建变量类型
-  bool string
- (u)int (u)int8 (u)int16,   (u)int32,(u)int64, uintptr 指针  加u无符号证书,不加u有符号整数,根据操作系统分,规定长度,不规定长度
- byte rune 字符型,go语言的char类型,byte 8位,rune 32位
- float32,float64,complex64,complex128 复数类型,complex64 的实部和虚部都是float32,complex128 实部和虚部都是float64

### 复数测试（欧拉公式）
![](/Users/guohewei/Learn/Golang/learnGo/chapter_02/image/complex_introduce.png）