package anyQueue

// 方法二 通过别名扩展系统包
type Queue []interface{}

// interface{} 任何类型都可以 像python的列表
func (q *Queue) Append(val interface{}) {
	*q = append(*q, val)
}

func (q *Queue) Pop() interface{} {
	top := (*q)[0]
	*q = (*q)[1:]
	return top
}

// 限制 interface{} 变量的类型 方法一
/*func (q *Queue) Append(val int) {
	*q = append(*q, val)
}

func (q *Queue) Pop() int {
	top := (*q)[0]
	*q = (*q)[1:]
	return top.(int)    // top.(int) 限制 top 类型
}*/

// 限制 interface{} 变量的类型 方法一
/*func (q *Queue) Append(val interface{}) {
	*q = append(*q, val.(int))
}

func (q *Queue) Pop() interface{} {
	top := (*q)[0]
	*q = (*q)[1:]
	return top.(int)    // top.(int) 限制 top 类型
}*/

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
