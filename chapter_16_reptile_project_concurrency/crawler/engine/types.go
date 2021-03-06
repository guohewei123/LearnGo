package engine

// Request 定义request结构体
type Request struct {
	Url        string
	ParserFunc func([]byte) ParserResult
}

// ParserResult 定义解析器返回结果结构体
type ParserResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Id      string
	Type    string
	Payload interface{}
}

// NilParser 定义一个空的解析器结果
func NilParser([]byte) ParserResult {
	return ParserResult{}
}