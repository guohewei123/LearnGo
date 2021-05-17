package main

import (
	"crawler_distributied/config"
	"crawler_distributied/persist"
	"crawler_distributied/rpcsupport"
	"flag"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

var host = flag.String("host", ":0", "the host for me to listen on")

func main() {
	flag.Parse()
	if *host == ":0" {
		fmt.Println("must specify a host")
		return
	}
	log.Fatal(serverRpc(*host, config.ElasticIndex))

}

func serverRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetURL(config.ElasticBaseUrl),
		elastic.SetSniff(false))
	if err != nil {
		return err
	}
	//log.Printf("Listen on %s", host)
	return rpcsupport.ServerRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
