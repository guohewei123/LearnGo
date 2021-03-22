package tree

import "fmt"

// 为结构体定义方法
func (node NodeTree) print() {
	fmt.Printf("%d ", node.Val)
}

// 为结构体定义， 设置val 的方法，
// 需要传入: &NodeTree 的地址, 原因是：go 是值传递，必须传入值的地址，才能真正修改入参
func (node *NodeTree) SetVal(val int) {
	if node == nil {
		fmt.Println("setting value to nil node. Ignored!")
		return
	}
	node.Val = val
}

// 为结构体定义中序遍历方法
/*func (node *NodeTree) OrderTraversal() {
	if node == nil {
		return
	}
	node.Left.OrderTraversal()
	node.print()
	node.Right.OrderTraversal()
}*/

// 使用函数来遍历二叉树
func (node *NodeTree) OrderTraversal() {
	node.TraverseFunc(
		func(n *NodeTree) {
			n.print()
		})
}

func (node *NodeTree) TraverseFunc(f func(tree *NodeTree)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

// 没有构造函数，我们可以使用工厂函数
func CreateNode(val int) *NodeTree {
	return &NodeTree{Val: val}
}


// 通过 channel 遍历树
func (node *NodeTree) TraverseWithChannel() chan *NodeTree {
	out := make(chan *NodeTree)  // 创建一个 tree类型 channel
	go func() {
		// 调用 TraverseFunc 遍历树，传入遍历执行函数
		node.TraverseFunc(func(n *NodeTree) {
			out <- n            // 添加遍历 node 到 channel
		})
		close(out)              // 关闭 channel
	}()
	return out
}