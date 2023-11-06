package rawInfo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/Fabricio2210/gofiber/elastic"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type LogInfo struct {
	UserName string `json:"userName"`
	Text     string `json:"text"`
	Hour     string `json:"hour"`
	LogDay   string `json:"logDay"`
}

func Search(query map[string]interface{}, page int, limit int) ([]LogInfo, float64, error) {
	es, err := elastic.ConnectElastic()

	// Define the mapping
	mapping := `{
		"properties": {
			"date": {
				"type": "date"
			},
			"text": {
				"type": "text"
			},
			"userName": {
				"type": "text"
			}
		}
	}`

	// Create the index with mapping
	reqCreateIndex := esapi.IndicesCreateRequest{
		Index: "logschemas",
		Body:  bytes.NewReader([]byte(mapping)),
	}

	resCreateIndex, err := reqCreateIndex.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error creating index: %s", err)
		return nil, 0, err
	}
	defer resCreateIndex.Body.Close()

	// Convert query map to JSON
	queryJSON, err := json.Marshal(query)
	if err != nil {
		log.Fatalf("Error marshaling query to JSON: %s", err)
		return nil, 0, err
	}

	// Perform the search
	fromValue := (page * limit) + 1
	reqSearch := esapi.SearchRequest{
		Index: []string{"logschemas"},
		Body:  strings.NewReader(string(queryJSON)),
		From:  &fromValue,
		Size:  &limit,
	}

	resSearch, err := reqSearch.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error executing search: %s", err)
		return nil, 0, err
	}
	defer resSearch.Body.Close()

	// Process the search response
	var result map[string]interface{}
	if err := json.NewDecoder(resSearch.Body).Decode(&result); err != nil {
		log.Fatalf("Error decoding search response: %s", err)
		return nil, 0, err
	}

	// Extract the hits from the result
	hits, ok := result["hits"].(map[string]interface{})
	if !ok {
		return nil, 0, fmt.Errorf("invalid search response format")
	}

	// Extract the search results and total count
	hitsArray, ok := hits["hits"].([]interface{})
	if !ok {
		return nil, 0, fmt.Errorf("invalid search response format")
	}

	// Process the search results
	var searchData []LogInfo
	for _, hit := range hitsArray {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			return nil, 0, fmt.Errorf("invalid search response format")
		}
		sourceData, ok := hitMap["_source"].(map[string]interface{})
		if !ok {
			return nil, 0, fmt.Errorf("invalid search response format")
		}
		// Check and convert fields safely
		userName, userNameOk := sourceData["userName"].(string)
		text, textOk := sourceData["text"].(string)
		hour, hourOk := sourceData["hour"].(string)
		logDay, logDayOk := sourceData["logDay"].(string)

		if !userNameOk || !textOk || !hourOk || !logDayOk {
			// Handle missing or invalid data gracefully
			continue // Skip this entry or take appropriate action
		}

		logInfo := LogInfo{
			UserName: userName,
			Text:     text,
			Hour:     hour,
			LogDay:   logDay,
		}
		searchData = append(searchData, logInfo)
	}

	// Extract the count
	count := hits["total"].(map[string]interface{})["value"].(float64)

	// Construct the final response
	return searchData, count, nil
}
