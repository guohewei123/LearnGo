package main

import (
	"downloader/infra"
	mocktesting "downloader/mock"
	"fmt"
	"time"
)

// 定义接口
// ?: Something that can "Get"
type retriever interface {
	Get(string2 string) string
}

// 返回类型为接口类型
func getRetriever() retriever {
	//return mocktesting.Retriever{}
	return infra.Retriever{}
}

// 查看接口变了内部是什么？ 方法：switch type
func switchInspect(r retriever) {
	fmt.Printf("%T  %v |分割| ", r, r)
	switch v := r.(type) {
	case *mocktesting.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case infra.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}

// 查看接口变了内部是什么？ 方法：type assertion
func assertionInspect(r retriever) {
	fmt.Printf("Input r: %T  %v |分割| ", r, r)
	if mockR, ok := r.(*mocktesting.Retriever); ok {
		fmt.Printf("mockR.Contents: %s |分割| ", mockR.Contents)
	} else {
		fmt.Printf("r is not a *mocktesting.Retriever  |分割| ")
	}

	if infraR, ok := r.(infra.Retriever); ok {
		fmt.Printf("infraR.TimeOut: %d |分割| ", infraR.TimeOut)
	} else {
		fmt.Println("r is not a infra.Retriever")
	}
}

func main() {

	// 1. 测试接口代码
	//var r retriever = getRetriever()
	//fmt.Println(r.Get("https://www.baidu.com"))

	// 2. 查看接口变量内部是什么？
	var r1, r2 retriever
	r1 = &mocktesting.Retriever{Contents: "this is a fake obj"}
	//r1 = mocktesting.Retriever{Contents: "this is a fake obj"}  // 指针接收者实现只能以指针方式使用；值接收者都可以

	r2 = infra.Retriever{UserAgent: "Mozilla/5.0", TimeOut: time.Second}
	r2 = &infra.Retriever{UserAgent: "Mozilla/5.0", TimeOut: time.Second} // 指针接收者实现只能以指针方式使用；值接收者都可以
	// 通过 switch type 查看
	switchInspect(r1) // *mocktesting.Retriever  &{this is a fake obj} |分割| Contents:  this is a fake obj
	switchInspect(r2) // infra.Retriever  {Mozilla/5.0 1s} |分割| UserAgent:  Mozilla/5.0
	// 通过 type assertion 查看
	assertionInspect(r1) // Input r: *mocktesting.Retriever  &{this is a fake obj} |分割| mockR.Contents: this is a fake obj |分割| r is not a infra.Retriever
	assertionInspect(r2) // Input r: infra.Retriever  {Mozilla/5.0 1s} |分割| r is not a *mocktesting.Retriever  |分割| infraR.TimeOut: 1000000000 |分割|
}
