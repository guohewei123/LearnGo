package mocktesting

type Retriever struct{
	Contents string
}

// 接口定义时使用，课程3
//func (Retriever) Get(url string) string {
//	return "Mock: " + url
//}

// 接口内部查看时使用，课程4
func (*Retriever) Get(url string) string {
	return "Mock: " + url
}
