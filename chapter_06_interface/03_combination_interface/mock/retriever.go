package mocktesting

import "fmt"

type Retriever struct{
	Contents string
}

func (r *Retriever) String() string {  // fmt.Stringer 接口
	return fmt.Sprintf("Retriever: {Contents=%s}", r.Contents)  // fmt.Println(r) 的时候会自动调用该方法
}

// 接口定义时使用，课程3
//func (Retriever) Get(url string) string {
//	return "Mock: " + url
//}

// 接口内部查看时使用，课程4
func (r *Retriever) Get(url string) string {
	return "Mock: " + r.Contents
}

func (r *Retriever) Post(url string, data map[string]string) string {
	fmt.Printf("Success Post data to: %s \n", url)
	r.Contents = data["contents"]
	return "ok"
}