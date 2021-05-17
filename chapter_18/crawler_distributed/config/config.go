package config

const (
	// Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	NilParser     = "NilParser"

	// ElasticSearch
	ElasticIndex = "profile_data"
	ElasticBaseUrl = "http://192.168.10.53:9200"

	// Rate limiting
	Qps = 20

	// Service host
	ItemSaverHost = ":1234"
	WorkerPort0 = ":9000"

	// service RPC
	ItemSaverRpc = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"
)
