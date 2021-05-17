package main

import (
	"crawler_distributied/config"
	"crawler_distributied/rpcsupport"
	"crawler_distributied/worker"
	"fmt"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServerRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewRpcClient(host)
	if err != nil {
		panic(err)
	}
	
	req := worker.Request{
		Url:    "https://www.zhenai.com/zhenghun/shanghai",
		Parser: worker.SerializedParser{
			Name: config.ParseCity,
		},
	}

	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	}else {
		fmt.Println(result)
	}



}