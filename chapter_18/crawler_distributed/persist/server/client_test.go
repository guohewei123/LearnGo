package main

import (
	"crawler_distributied/config"
	"crawler_distributied/engine"
	"crawler_distributied/model"
	"crawler_distributied/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	// start ItemSaverServer
	go serverRpc(host, "test1")
	// start ItemSaverClient

	time.Sleep(time.Second)
	client, err := rpcsupport.NewRpcClient(host)
	if err != nil {
		panic(err)
	}

	expected := engine.Item{
		Url:     "https://album.zhenai.com/u/1876503328",
		Id:      "1876503328",
		Type:    "zhenai",
		Payload: model.Profile{
			Name:              "张三",
			Gender:            "男",
			Residence:         "北京",
			Age:               18,
			IncomeOrEducation: "3000-4000",
			Marriage:          "未婚",
			Height:            180,
		},
	}

	result := ""
	err = client.Call(config.ItemSaverRpc, expected, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %v", result, err)
	}

	// Call save
}
