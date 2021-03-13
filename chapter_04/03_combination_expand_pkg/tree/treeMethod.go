package tree

import "fmt"

// 为结构体定义方法
func (node NodeTree) Print() {
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
func (node *NodeTree) OrderTraversal() {
	if node == nil {
		return
	}
	node.Left.OrderTraversal()
	node.Print()
	node.Right.OrderTraversal()
}

// 没有构造函数，我们可以使用工厂函数
func CreateNode(val int) *NodeTree {
	return &NodeTree{Val: val}
}
