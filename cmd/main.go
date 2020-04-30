package main

import (
	"elasticstudy.com/elastic/pkg/client"
	"elasticstudy.com/elastic/pkg/service"
	"elasticstudy.com/elastic/pkg/types"
	"log"
	"time"
)

func main() {

	clientES, err := client.GetESClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// INSERT DOCUMENT
	newDocument := types.Document{
		Name:        "Document",
		Description: "This document is a test!",
		Score:       55.5,
		Create:      time.Now(),
	}

	service.Indexer(newDocument, clientES)

	// SEARCH WITH INDEX
	service.SearchById("123456789", clientES)

	// SEARCH DOCUMENT
	service.Search(clientES)

	// DELETE DOCUMENT
	service.Delete("123456789", clientES)

}
