package main

import (
	"fmt"
	"treeTesting/tree"
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
	root.Right.Left = &tree.NodeTree{Val: 4}
	root.OrderTraversal() // 0 2 3 4 5
	fmt.Println()

	treeCount := 0
	root.TraverseFunc(func(node *tree.NodeTree) {
		treeCount += 1
	})
	fmt.Printf("treeCount: %d\n", treeCount)

	// 使用 channel 遍历树
	c := root.TraverseWithChannel()
	maxNode := 0
	for n := range c {   // 获取 channel 中的 node
		if n.Val > maxNode {
			maxNode = n.Val
		}
	}
	fmt.Println("max Node Value: ", maxNode)

}