package service

import (
	"bytes"
	"context"
	"elasticstudy.com/elastic/pkg/types"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
)

func Update(id string, document types.Document, clientES *elasticsearch.Client) {

	var doc bytes.Buffer
	documentUp := map[string]interface{}{
		"doc": map[string]interface{}{
			"name": document.Name,
		},
	}

	if err := json.NewEncoder(&doc).Encode(documentUp); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	 req := esapi.UpdateRequest{
		 Index:               "documents",
		 DocumentID:          id,
		 Body:                &doc,
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
		var r map[string]interface{}
		if errJs := json.NewDecoder(res.Body).Decode(&r); errJs != nil {
			log.Printf("Error parsing the response body: %s", errJs)
		} else {
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
}
