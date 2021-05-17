package engine

type Parser interface {
	Parser(contents []byte) ParserResult
	Serialize() (name string, args interface{})
}

type ParserFunc func(contents []byte) ParserResult

// 定义request结构体
type Request struct {
	Url    string
	Parser Parser
}

// 定义解析器返回结果结构体
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

// 定义一个空的解析器结果

type NilParser struct{}

func (NilParser) Parser(_ []byte) ParserResult {
	return ParserResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

// -----------------
type FuncParser struct {
	parser ParserFunc
	Name   string
}

func (f *FuncParser) Parser(contents []byte) ParserResult {
	return f.parser(contents)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.Name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		Name:   name,
	}
}
