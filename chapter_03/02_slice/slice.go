package main

import "fmt"

func printSlice(sliceName string, s []int) {
	fmt.Printf("Slice %s=%v, len=%d, cap=%d\n", sliceName, s, len(s), cap(s))
}

func appendSlice() {
	var s []int
	for i := 0; i < 10; i++ {
		printSlice("s", s)
		s = append(s, i*2)
		s = append(s, i*2+1)
	}
}

func main() {
	arr := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	// Slice本身没有数据，是对底层 array 的一个 view
	fmt.Println("------------定义切片----------")
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])

	// 切片的切片 Re_slice(Slice the Slice)
	fmt.Println("------------切片的切片----------")
	printSlice("arr", arr[:])
	s1 := arr[2:6]
	printSlice("s1", s1)
	s2 := s1[3:5]
	printSlice("s2", s2)

	// 切片超出 cap, 将会报错：panic: runtime error: slice bounds out of range [:7] with capacity 6
	fmt.Println("------------切片超出 cap----------")
	printSlice("arr", arr[:])
	s1 = arr[2:6]
	printSlice("s1", s1)
	//s2 = s1[3:7]
	//printSlice("s2", s2)

	// 切片添加元素 append
	fmt.Println("------------切片添加元素 append----------")
	printSlice("arr", arr[:])
	s1 = arr[2:7]
	printSlice("s1", s1)
	s1 = append(s1, 10)
	printSlice("s1", s1)
	printSlice("arr", arr[:])
	// 超出 arr 的 cap, 系统将会定义一个新的arr来对应 s1, 旧的arr在系统没有使用的情况下会被垃圾回收
	s1 = append(s1, 11)
	printSlice("s1", s1)
	printSlice("arr", arr[:])

	// 使用 make 定义切片 （可以指定切片的 len 和 cap）
	fmt.Println("------------使用 make 定义切片----------")
	makeS1 := make([]int, 10, 32) // Slice make_s1=[0 0 0 0 0 0 0 0 0 0], len=10, cap=32
	printSlice("make_s1", makeS1)

	// 拷贝slice到另一个slice
	fmt.Println("------------拷贝slice到另一个slice----------")
	arr = [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 = arr[2:7]
	s2 = make([]int, 10, 16)
	printSlice("s1", s1)
	printSlice("s2", s2)
	num := copy(s2, s1) // func copy(dst, src []Type) int
	fmt.Println("拷贝 s1 到另一个 s2, 总计 copy 个数: ", num)
	printSlice("s1", s1)
	printSlice("s2", s2)

	// 删除 slice 中的元素, 没有内建方法，需要通过 copy 覆盖掉要删除的元素
	fmt.Println("------------ 删除 slice 中 第3个元素 ----------")
	srcS1 := []int{0, 1, 2, 3, 4, 5, 6, 7}
	printSlice("srcS1", srcS1)
	srcS1 = append(srcS1[:3], srcS1[4:]...)
	printSlice("srcS1", srcS1)

	// Pop slice 中的第一元素, Pop slice 中的最后一元素
	fmt.Println("------------ Pop slice 中的第一元素, Pop slice 中的最后一元素 ----------")
	fmt.Println("Pop slice 中的第一元素")
	srcS1 = []int{0, 1, 2, 3, 4, 5, 6, 7}
	printSlice("srcS1", srcS1)
	fmt.Println("Pop slice first element: ", srcS1[0])
	srcS1 = srcS1[1:]
	printSlice("srcS1", srcS1)

	fmt.Println("Pop slice last element: ", srcS1[len(srcS1)-1])
	srcS1 = srcS1[:len(srcS1)-1]
	printSlice("srcS1", srcS1)

	// 测试 slice 的自动扩容
	fmt.Println("------------ 测试 slice 的自动扩容 ----------")
	appendSlice()
}
