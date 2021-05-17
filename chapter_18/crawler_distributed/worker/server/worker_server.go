package main

import (
	"crawler_distributied/rpcsupport"
	"crawler_distributied/worker"
	"flag"
	"fmt"
	"log"
)

var host = flag.String("host", ":0", "the host for me to listen on")

func main() {
	flag.Parse()
	if *host == ":0" {
		fmt.Println("must specify a host")
		return
	}
	log.Fatal(rpcsupport.ServerRpc(*host, worker.CrawlService{}))
}
