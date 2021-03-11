package main

import (
	"fmt"
	"maketree/tree"
)

func main() {
	fmt.Println("---- 分割线 ---")
	// 2. 定义tree
	//     3
	//   /  \
	//  0    5
	//  \    /
	//   2  4
	var root = tree.NodeTree{Val: 3}
	root.Left = new(tree.NodeTree)
	root.Right = &tree.NodeTree{Val: 5}
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left = &tree.NodeTree{Val:4}
	root.OrderTraversal() // 0 2 3 4 5
	fmt.Println()
}
