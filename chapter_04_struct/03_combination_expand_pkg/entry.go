package main

import (
	"fmt"
	"maketree/tree"
)

// 方法一 通过组合扩展包
type myTree struct {
	node *tree.NodeTree
}

func (myNode *myTree) postTraversal() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left :=myTree{myNode.node.Left}
	left.postTraversal()
	right := myTree{myNode.node.Right}
	right.postTraversal()
	myNode.node.Print()
}


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
	fmt.Print("OrderTraversal: ")
	root.OrderTraversal() // 0 2 3 4 5
	fmt.Println()

	myRoot := myTree{&root}
	fmt.Print("postTraversal: ")
	myRoot.postTraversal()
}
