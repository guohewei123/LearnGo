# 内建容器
## 一、数组
- 数组定义
    ```
    var arr1 [5]int // 声明数组
    arr2 := [3]int{1, 3, 5}  // 声明数组并赋值
    arr3 := [...]int{2, 4, 6, 8, 10} // 不输入数组长度，让编译器来计算长度
    var grid [4][5]int // 二维数组
    ```
 
 - 数量写在类型前
 - 可通过 _ 来省略变量，不仅仅是 range，任何地方都可通过 _ 来省略变量
    ```
    sum := 0
    for _, v := range numbers {
        sum += v
    }
    ```
 - 如果只要下标 i，可写成for i := range numbers
 - 数组是值类型
    - [10]int 和[20]int 是不同类型
    - 调用 func f(arr [10]int)会 拷贝 数组
    -  在 go 语言中一般不直接使用数组（指针），使用切片
 
```go
package main

import "fmt"

func printArray(arr [5]int) {
	fmt.Println("Traversal i")
	for i := range arr {
		fmt.Println(arr[i])
	}
	fmt.Println("Traversal (i, val)")
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Println("Traversal (_, val)")
	for _, v := range arr {
		fmt.Println(v)
	}
}

// 值类型传递 函数中修改入参数组，函数外不会被修改
func modifyArray1(arr [5]int) {
	arr[0] = 100
	fmt.Println(arr)
}

// 值类型传递 函数中修改入参数组，函数外不会被修改
func modifyArray2(arr *[5]int) {
	(*arr)[0] = 100
	arr[1] = 101
	fmt.Println(*arr)
}

func main() {
	fmt.Printf("-------------Define arr test-----------\n")
	var arr1 [5]int
	var arr2 = [5]int{2, 4, 6, 7, 8}
	arr3 := [3]int{1, 3, 5}
	arr4 := [...]int{2, 4, 6, 7, 8}
	var arr5 [3][5]int

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)

	fmt.Printf("-------------Print arr2 test: %d-------------\n", arr4)
	printArray(arr2)

	fmt.Printf("-------------Modify arr4 test: %d-------------\n", arr4)
	modifyArray1(arr4)
	fmt.Println(arr4)
	modifyArray2(&arr4)
	fmt.Println(arr4)
}
```
输出结果：
```
-------------Define arr test-----------
[0 0 0 0 0]
[2 4 6 7 8]
[1 3 5]
[2 4 6 7 8]
[[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]
-------------Print arr2 test: [2 4 6 7 8]-------------
Traversal i
2
4
6
7
8
Traversal (i, val)
0 2
1 4
2 6
3 7
4 8
Traversal (_, val)
2
4
6
7
8
-------------Modify arr4 test: [2 4 6 7 8]-------------
[100 4 6 7 8]
[2 4 6 7 8]
[100 101 6 7 8]
[100 101 6 7 8]

```
