package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemServer("profile_data")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		//Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 100,
		ItemChan: itemChan,
	}
	e.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	//e.Run(engine.Request{
	//	Url:        "https://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}
