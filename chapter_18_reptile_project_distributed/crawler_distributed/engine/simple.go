package engine

import (
	"log"
)

type SimpleEngine struct {}

func (SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parserResult, err := Worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parserResult.Requests...)  // 将解析后需要继续request的添加到队列中
		for _, item := range parserResult.Items {
			log.Printf("Got item %v\n", item)                  // 打印获取到的结果
		}
	}
}
