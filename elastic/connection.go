package elastic

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

func ConnectElastic() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		// Log the error and return it
		log.Printf("Failed to create Elasticsearch client: %s", err)
	}

	// Test the connection
	res, err := client.Info()
	if err != nil {
		// Log the error
		log.Printf("Failed to retrieve Elasticsearch info: %s", err)
	} else {
		// Check the response status
		if res.IsError() {
			log.Printf("Error response: %s", res.String())
		}

		// Process the response
		fmt.Println(res.String())
	}

	return client, nil
}
