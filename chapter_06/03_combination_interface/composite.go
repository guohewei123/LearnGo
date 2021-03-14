package main

import (
	mocktesting "compositeIntreface/mock"
	"fmt"
)

// 定义接口 retriever
type retriever interface {
	Get(string2 string) string
}

// 定义接口 post
type submit interface {
	Post(string, map[string]string) string
}

// 定义组合接口 retriever and submit
type retrieverSubmit interface {
	retriever
	submit
}

func session(r retrieverSubmit) string {

	r.Post("https://www.baidu.com", map[string]string{"contents": "Another fake web"})
	return r.Get("https://www.baidu.com")
}

func main() {
	// 1. 测试接口代码
	r := mocktesting.Retriever{Contents:"This is a fake web"}
	//fmt.Println(session(&r))
	fmt.Println(&r)   // Retriever: {Contents=Another fake web}

}
