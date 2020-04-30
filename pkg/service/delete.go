package service

import (
	"context"
	"elasticstudy.com/elastic/pkg/types"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
)

func Delete(id string, clientES *elasticsearch.Client) {
	req := esapi.DeleteRequest{
		Index:               "documents",
		DocumentID:          id,
		Refresh:             "true",
		Pretty:              true,
	}

	res, err := req.Do(context.Background(), clientES)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error delete document", res.Status())
	} else {
		var r types.ResponseES
		if errJs := json.NewDecoder(res.Body).Decode(&r); errJs != nil {
			log.Printf("Error parsing the response body: %s", errJs)
		} else {
			log.Printf("[%s] version=%d ID=%s Index=%s", res.Status(), r.Version, r.ID, r.Index)
		}
	}
}