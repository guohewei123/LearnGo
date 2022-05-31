package main

import (
	"crawler_distributied/config"
	"crawler_distributied/engine"
	itemServer "crawler_distributied/persist/client"
	"crawler_distributied/rpcsupport"
	"crawler_distributied/scheduler"
	worker "crawler_distributied/worker/client"
	"crawler_distributied/zhenai/parser"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	//itemChan, err := persist.ItemServer(config.ElasticIndex)
	itemChan, err := itemServer.ItemServer(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		//Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
		//RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url:    "https://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})

	//e.Run(engine.Request{
	//	Url:        "https://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewRpcClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
