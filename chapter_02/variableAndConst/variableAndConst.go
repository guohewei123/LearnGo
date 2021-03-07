package main

import (
	"fmt"
	"math"
	"math/cmplx"
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

// 欧拉公式

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i * math.Pi) + 1)
}

// 类型转换
func triangle() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 常量定义
func consts() {
	const filename = "abc.txt"
	const (
		a, b = 3, 4
		name = "lao wang"
	)
	c := math.Sqrt(a*a + b*b)
	fmt.Println(filename, c, name)
}

// 使用常量定义枚举类型
func enums() {
	//普通枚举类型
	/*const(
		cpp = 0
		java = 1
		python = 2
		golang = 3
		javascript = 4
	)*/
	//自增值枚举类型
	const(
		cpp = iota
		java
		python
		golang
		javascript
	)
	fmt.Println(cpp, java, python, golang, javascript)

	// b, kb, mb, gb, tb, pb
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}


func main() {
	variableZeroValue()
	variableInitValue()
	varTypeDefvalue()
	variableValueShort()
	fmt.Println(aa, bb, cc)
	euler()
	triangle()
	consts()
	enums()
}
