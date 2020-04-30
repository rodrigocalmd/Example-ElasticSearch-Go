package service

import (
	"context"
	"elasticstudy.com/elastic/pkg/types"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strings"
)

func Indexer(document types.Document, clientES *elasticsearch.Client) {
	dataJSON, err := json.Marshal(document)
	js := string(dataJSON)

	req := esapi.IndexRequest{
		Index: "documents",
		Body:	strings.NewReader(js),
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), clientES)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document", res.Status())
	} else {
		var r types.ResponseES
		if errJs := json.NewDecoder(res.Body).Decode(&r); errJs != nil {
			log.Printf("Error parsing the response body: %s", errJs)
		} else {
			log.Printf("[%s] version=%d ID=%s Index=%s", res.Status(), r.Version, r.ID, r.Index)
		}
	}
}
