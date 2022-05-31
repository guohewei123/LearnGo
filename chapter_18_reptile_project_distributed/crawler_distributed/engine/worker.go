package engine

import (
	"crawler_distributied/fetcher"
	"log"
)

func Worker(r Request) (ParserResult, error) {
	//log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)    // 获取数据
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParserResult{}, err
	}
	return r.Parser.Parser(body), nil // 调用指定的解析器解析数据
	//return r.ParserFunc(body), nil
}