package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// 通过指针来交换值
func swap1(a, b *int) {
	*a, *b = *b, *a
}

func swap2(a, b int) (int, int) {
	return b, a
}

func swapValueOfVariable() {
	a, b := 3, 4
	swap1(&a, &b)
	fmt.Println(a, b)
	a, b = 3, 4
	a, b = swap2(a, b)
	fmt.Println(a, b)
}

// 获取当前代码所在路径
func runtimeCaller() {
	_, filename, _, ok := runtime.Caller(1)
	var cwdPath string
	if ok {
		cwdPath = path.Join(path.Dir(filename), "") // the the main function file directory
	} else {
		cwdPath = "./"
	}
	fmt.Println("cwd path...", cwdPath)
}

// 获取当前执行程序所在的绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}


func main() {

	swapValueOfVariable()
	runtimeCaller()
	fmt.Println("getCurrentAbPathByExecutable = ", getCurrentAbPathByExecutable())

	// NumCPU returns the number of logical
	// CPUs usable by the current process.
	fmt.Println(runtime.NumCPU())
}
