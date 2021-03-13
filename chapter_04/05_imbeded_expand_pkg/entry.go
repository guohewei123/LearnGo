package main

import (
	"fmt"
	"maketree/tree"
)

// 方法三 通过内嵌方式扩展包  embedding
type myTree struct {
	*tree.NodeTree  // 语法糖 embedding
}

// 扩展 postTraversal
func (myNode *myTree) postTraversal() {
	if myNode == nil || myNode.NodeTree == nil {
		return
	}
	left := myTree{myNode.Left}
	right := myTree{myNode.Right}
	left.postTraversal()
	right.postTraversal()
	myNode.Print()
}

// 定义与NodeTree同名的方法： 将底层的NodeTree中OrderTraversal将被shadowed, 可通过 node.NodeTree.OrderTraversal 调用
func (myNode *myTree) OrderTraversal() {
	if myNode == nil || myNode.NodeTree == nil {
		return
	}
	left := myTree{myNode.Left}
	right := myTree{myNode.Right}
	left.OrderTraversal()
	fmt.Print("myTree: ")
	myNode.Print()
	right.OrderTraversal()
}


func main() {
	fmt.Println("---- 分割线 ---")
	// 2. 定义tree
	//     3
	//   /  \
	//  0    5
	//  \    /
	//   2  4
	var root = myTree{&tree.NodeTree{Val:3}}
	root.Left = new(tree.NodeTree)
	root.Right = &tree.NodeTree{Val: 5}
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left = &tree.NodeTree{Val:4}

	//myRoot := myTree{&root}
	fmt.Print("Expend method postTraversal     : ")
	root.postTraversal()            // Expend method postTraversal     : 2 0 4 5 3

	fmt.Println()
	fmt.Print("myTree orderTraversal           : ")
	root.OrderTraversal()           // myTree orderTraversal           : myTree: 0 myTree: 2 myTree: 3 myTree: 4 myTree: 5

	fmt.Println()
	fmt.Print("Shadowed NodeTree orderTraversal: ")
	root.NodeTree.OrderTraversal()  // Shadowed NodeTree orderTraversal: NodeTree: 0 NodeTree: 2 NodeTree: 3 NodeTree: 4 NodeTree: 5
}
