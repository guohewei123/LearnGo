package engine

// 定义request结构体
type Request struct {
	Url        string
	ParserFunc func([]byte) ParserResult
}

// 定义解析器返回结果结构体
type ParserResult struct {
	Requests []Request
	Items    []interface{}
}

// 定义一个空的解析器结果
func NilParser([]byte) ParserResult {
	return ParserResult{}
}