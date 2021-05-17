package persist

import (
	"context"
	"crawler_distributied/config"
	"crawler_distributied/engine"
	"errors"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemServer(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	client, err := elastic.NewClient(
		elastic.SetURL(config.ElasticBaseUrl),
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			err := save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v:%v", item, err)
				continue
			}
			itemCount++
		}

	}()
	return out, nil
}

func save(client *elastic.Client, index string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must support type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())

	if err != nil {
		return err
	}
	return nil
	//fmt.Printf("%+v\n", resp)
}
