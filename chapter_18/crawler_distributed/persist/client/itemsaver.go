package client

import (
	"crawler_distributied/config"
	"crawler_distributied/engine"
	"crawler_distributied/rpcsupport"
	"log"
)

func ItemServer(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewRpcClient(host)
	if err != nil {
		panic(err)
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			// Call RPC to save item
			result := ""
			err = client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Item Saver: error saving item %v:%v", item, err)
				continue
			}
			itemCount++
		}

	}()
	return out, nil
}
