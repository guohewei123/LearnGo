package main

import (
	"fmt"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"
)

func foo() {
}

func GetFunctionName(i interface{}, seps ...rune) string {
	// 获取函数名称
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()

	// 用 seps 进行分割
	fields := strings.FieldsFunc(fn, func(sep rune) bool {
		for _, s := range seps {
			if sep == s {
				return true
			}
		}
		return false
	})

	// fmt.Println(fields)

	if size := len(fields); size > 0 {
		return fields[size-1]
	}
	return ""
}

func main() {
	// This will print "name: main.foo"
	fmt.Println("name:", GetFunctionName(foo))

	// runtime/debug.FreeOSMemory
	fmt.Println(GetFunctionName(debug.FreeOSMemory))
	// FreeOSMemory
	fmt.Println(GetFunctionName(debug.FreeOSMemory, '.'))
	// FreeOSMemory
	fmt.Println(GetFunctionName(debug.FreeOSMemory, '/', '.'))
}