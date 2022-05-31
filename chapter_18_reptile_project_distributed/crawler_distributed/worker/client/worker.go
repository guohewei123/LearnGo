package client

import (
	"crawler_distributied/config"
	"crawler_distributied/engine"
	"crawler_distributied/worker"
	"net/rpc"
)

/*func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewRpcClient(config.WorkerPort0)
	if err != nil {
		return nil, err
	}
	return func(req engine.Request) (engine.ParserResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult

		err = client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParserResult{}, err
		}

		return worker.DeserializeParseResult(sResult), nil
	}, nil
}*/


func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(req engine.Request) (engine.ParserResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParserResult{}, err
		}
		return worker.DeserializeParseResult(sResult), nil
	}
}