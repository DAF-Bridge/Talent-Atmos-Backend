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

// type SearchQuery struct {
// 	Query struct {
// 		Match struct {
// 			Name string `json:"name"`
// 		} `json:"match"`
// 	} `json:"query"`
// }

func SearchEvents(client *opensearch.Client, query models.SearchQuery, page int, perPage int) ([]map[string]interface{}, error) {
	// query := SearchQuery{}
	// query.Query.Match.Name = keyword

	searchQuery := buildSearchQuery(query)

	queryBody, _ := json.Marshal(searchQuery)

	res, err := client.Search(
		client.Search.WithIndex("events"),
		client.Search.WithBody(bytes.NewReader(queryBody)),
		client.Search.WithContext(context.Background()),
	)
	if err != nil {
		logs.Error(fmt.Sprintf("failed to execute search: %v", err))
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

	return results, nil
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
				"fields": []string{"name", "description", "key_takeaway", "highlight", "location_name"},
				"type":   "best_fields", // Can be changed to "most_fields" / "cross_fields" / "phrase" / "phrase_prefix" for optimization
			},
		})
	}
	if query.Location != "" {
		must = append(must, map[string]interface{}{
			"match": map[string]interface{}{
				"location_type": query.Location,
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
				"price_type": query.PriceType,
			},
		})
	}
	if !startDate.IsZero() && !endDate.IsZero() {
		must = append(must, map[string]interface{}{
			"range": map[string]interface{}{
				"start_date": map[string]interface{}{
					"gte": startDate,
					"lte": endDate,
				},
			},
		})
	}

	if query.Page != 0 {
		searchQuery["from"] = query.Page
	}
	if query.PerPage != 0 {
		searchQuery["size"] = query.PerPage
	}

	// Add other filters similarly...

	boolQuery["must"] = must
	searchQuery["query"] = map[string]interface{}{
		"bool": boolQuery,
	}

	return searchQuery
}
