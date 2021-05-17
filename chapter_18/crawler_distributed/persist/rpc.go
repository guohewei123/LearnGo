package persist

import (
	"crawler_distributied/engine"
	"github.com/olivere/elastic/v7"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := save(s.Client, s.Index, item)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		*result = "failed"
		log.Printf("Error saving item %v: %v", item, err)
	}
	return err
}
