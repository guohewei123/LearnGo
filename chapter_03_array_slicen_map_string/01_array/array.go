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
