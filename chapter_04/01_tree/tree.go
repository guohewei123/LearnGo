package main

import "fmt"

type nodeTree struct {
	val         int
	left, right *nodeTree
}

// 为结构体定义方法
func (node nodeTree) print() {
	fmt.Printf("%d ", node.val)
}

// 为结构体定义， 设置val 的方法，
// 需要传入: &nodeTree 的地址, 原因是：go 是值传递，必须传入值的地址，才能真正修改入参
func (node *nodeTree) setVal(val int) {
	if node == nil {
		fmt.Println("setting value to nil node. Ignored!")
		return
	}
	node.val = val
}

// 为结构体定义中序遍历方法
func (node *nodeTree) orderTraversal() {
	if node == nil {
		return
	}
	node.left.orderTraversal()
	node.print()
	node.right.orderTraversal()
}

// 没有构造函数，我们可以使用工厂函数
func createNode(val int) *nodeTree {
	return &nodeTree{val: val}
}

// 结构体变量的定义
func testStructDef() {

	// 1. 定义结构体变量
	var root = nodeTree{val: 3}                       // 方法一
	root.left = new(nodeTree)                         // 方法二
	root.right = &nodeTree{}                          // 方法二
	root.left.right = &nodeTree{val: 0}               // 方法三
	root.right.left = createNode(2)                   // 方法四
	fmt.Println("root: ", root)                       // root:  {3 0xc00000c080 0xc00000c0a0}
	fmt.Println("root.left: ", root.left)             // root.left:  &{0 <nil> 0xc00000c0c0}
	fmt.Println("root.right: ", root.right)           // root.right:  &{0 0xc00000c0e0 <nil>}
	fmt.Println("root.left.right: ", root.left.right) // root.left.right:  &{0 <nil> <nil>}
	fmt.Println("root.right.left: ", root.right.left) // root.right.left:  &{2 <nil> <nil>}

	// 2. 定义结构体Slice
	nodes := []nodeTree{
		{val: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes) // [{3 <nil> <nil>} {0 <nil> <nil>} {6 <nil> 0xc00000c060}]
}

func main() {

	// 1. 结构体变量的定义
	testStructDef()
	fmt.Println("---- 分割线 ---")
	// 2. 定义tree
	//     3
	//   /  \
	//  0    5
	//  \    /
	//   2  4
	var root = nodeTree{val: 3}
	root.left = new(nodeTree)
	root.right = &nodeTree{val: 5}
	root.left.right = createNode(2)
	root.right.left = &nodeTree{4, nil, nil}
	root.orderTraversal() // 0 2 3 4 5
	fmt.Println()
}
