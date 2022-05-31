package engine

import (
	"crawler/fetcher"
	"fmt"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)    // 获取数据
		if err != nil {
			fmt.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}
		parserResult := r.ParserFunc(body)  // 调用指定的解析器解析数据
		requests = append(requests, parserResult.Requests...)  // 将解析后需要继续request的添加到队列中
		for _, item := range parserResult.Items {
			fmt.Printf("Got item %v\n", item)                  // 打印获取到的结果
		}
	}
}
