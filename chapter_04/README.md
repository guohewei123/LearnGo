# 面向 "对象"
- go 语言仅支持封装，不支持继承和多态
- go语言没有 class，只有 struct
## 结构体和方法
- 不论地址还是结构本身，一律使用 . 来访问成员
- 结构创建在堆上还是栈上？ 不需要知道 自动垃圾回收
- 可以使用指针作为方法的接收者
    - 只有使用指针才可以改变结构体内容
    - nil 指针也可以调用方法
    
- 值接收者 vs 指针接收者
  - 要改变内容必须使用指针接收者
  - 结构体过大也考虑使用指针接收者（性能考虑）
  - 一致性：如有指针接收者，最好都是指针接收者
  - 值接收者是 go语言特有
  - 值/指针接收者的方法都是使用 . 来调用
- 结构体测试代码
    ```go
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
    ```
 
 ## 2. 包和封装
### 1. 封装
- 名字一般使用 CamelCase
- 首字母大写：public
- 首字母小写：private
### 2. 包
- 每个目录一个包
- main 包包含可执行入口
- 为结构定义的方法必须放在同一包内, 可以是不同文件

- 测试包的使用（详见代码）
    ```
    02_tree_pkg    
    ├── entry.go   // main 包
    ├── go.mod
    └── tree       // node tree 包
        ├── tree.go
        └── treeMethod.go
    ```

## 3. 如何扩充系统类型或者别人的类型
### 1. 使用组合
- 使用方法详见示例代码
### 2. 定义别名
- 使用方法详见示例代码
### 3. 内嵌方式
- 使用方法详见示例代码

 