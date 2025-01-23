package search

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"github.com/opensearch-project/opensearch-go"
)

func SearchEvents(client *opensearch.Client, query models.SearchQuery, page int, Offset int) (map[string]interface{}, error) {
	searchQuery := buildSearchQuery(query)

	queryBody, _ := json.Marshal(searchQuery)

	res, err := client.Search(
		client.Search.WithIndex("events"),
		client.Search.WithBody(bytes.NewReader(queryBody)),
		client.Search.WithContext(context.Background()),
	)
	if err != nil {
		logs.Error(fmt.Sprintf("failed to execute search on: %v", err))
		return nil, err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	hits, ok := result["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		logs.Error("No search results found")
		return nil, nil
	}

	var results []map[string]interface{}
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		results = append(results, source)
	}

	var responses map[string]interface{}

	// if there does not exist data in the hits, then we can return an empty response
	if len(hits) == 0 {
		responses = map[string]interface{}{
			"total_events": 0,
			"events":       []map[string]interface{}{},
		}
		return responses, nil
	}

	// if there exist data in the hits, then we can get the total hits
	totalHits := result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)
	responses = map[string]interface{}{
		"total_events": totalHits,
		"events":       results,
	}

	return responses, nil
}

func buildSearchQuery(query models.SearchQuery) map[string]interface{} {
	startDate, endDate := utils.GetDateRange(query.DateRange)

	// Construct the query map based on the filters
	searchQuery := make(map[string]interface{})
	boolQuery := make(map[string]interface{})
	must := []map[string]interface{}{}

	if query.Search != "" {
		must = append(must, map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  query.Search,
				"fields": []string{"name", "description", "keyTakeaway", "highlight", "location"},
				"type":   "best_fields", // Can be changed to "most_fields" / "cross_fields" / "phrase" / "phrase_prefix" for optimization
			},
		})
	}
	if query.Location != "" {
		must = append(must, map[string]interface{}{
			"match": map[string]interface{}{
				"locationType": query.Location,
			},
		})
	}
	if query.Category != "" {
		must = append(must, map[string]interface{}{
			"match": map[string]interface{}{
				"category": query.Category,
			},
		})
	}
	if query.Audience != "" {
		must = append(must, map[string]interface{}{
			"match": map[string]interface{}{
				"audience": query.Audience,
			},
		})
	}
	if query.PriceType != "" {
		must = append(must, map[string]interface{}{
			"match": map[string]interface{}{
				"priceType": query.PriceType,
			},
		})
	}
	if !startDate.IsZero() && !endDate.IsZero() {
		must = append(must, map[string]interface{}{
			"range": map[string]interface{}{
				"startDate": map[string]interface{}{
					"gte": startDate,
					"lte": endDate,
				},
			},
		})
	}
	// if query.Page != 0 {
	// 	searchQuery["from"] = query.Page
	// }
	if query.Page != 0 {
		searchQuery["from"] = (query.Page - 1) * query.Offset
	}
	if query.Offset != 0 {
		searchQuery["size"] = query.Offset
	}

	// Add other filters similarly...

	boolQuery["must"] = must
	searchQuery["query"] = map[string]interface{}{
		"bool": boolQuery,
	}

	return searchQuery
}
