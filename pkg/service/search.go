package service

import (
	"bytes"
	"context"
	"elasticstudy.com/elastic/pkg/types"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"strings"
)

func SearchById(id string, clientES *elasticsearch.Client) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"_id": id,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := clientES.Search(
		clientES.Search.WithContext(context.Background()),
		clientES.Search.WithIndex("documents"),
		clientES.Search.WithBody(&buf),
		clientES.Search.WithTrackTotalHits(true),
		clientES.Search.WithPretty(),
		)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	var r types.SearchResponseES
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	log.Printf(
		"[%s] %d hits; took: %dms", res.Status(), r.Hits.Total.Value, r.Took,
	)

	for _, hit := range r.Hits.Hits {
		log.Printf(" * ID=%s, %s", hit.ID, hit.Source)
	}

	log.Println(strings.Repeat("=", 37))
}

func Search(clientES *elasticsearch.Client) {

	res, err := clientES.Search(
		clientES.Search.WithContext(context.Background()),
		clientES.Search.WithIndex("documents"),
		clientES.Search.WithTrackTotalHits(true),
		clientES.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	var r types.SearchResponseES
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms", res.Status(), r.Hits.Total.Value, r.Took,
	)
	// Print the ID and document source for each hit.
	for _, hit := range r.Hits.Hits {
		log.Printf(" * ID=%s, %s", hit.ID, hit.Source)
	}

	log.Println(strings.Repeat("=", 40))
}

