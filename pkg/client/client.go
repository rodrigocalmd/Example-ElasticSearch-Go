package client

import (
	"github.com/elastic/go-elasticsearch/v8"
	"log"
)

func GetESClient() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	res, errIn := es.Info()
	if errIn != nil {
		log.Fatalf("Error getting response: %s", errIn)
	}
	log.Println(res)
	return es, err
}