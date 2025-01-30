package search

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"github.com/opensearch-project/opensearch-go"
)

func SearchEvents(client *opensearch.Client, query models.SearchQuery, page int, offset int) (dto.SearchEventResponse, error) {
	searchQuery := buildSearchQuery(query)

	queryBody, err := json.Marshal(searchQuery)
	if err != nil {
		logs.Error(fmt.Sprintf("failed to marshal search query: %v", err))
		return dto.SearchEventResponse{}, err
	}

	res, err := client.Search(
		client.Search.WithIndex("events"),
		client.Search.WithBody(bytes.NewReader(queryBody)),
		client.Search.WithContext(context.Background()),
	)
	if err != nil {
		logs.Error(fmt.Sprintf("failed to execute search on: %v", err))
		// return nil, err
		return dto.SearchEventResponse{}, err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	hits, ok := result["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		logs.Error("No search results found")
		// return nil,
		return dto.SearchEventResponse{}, nil
	}

	var results []models.EventDocument
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		event := models.EventDocument{}
		jsonString, _ := json.Marshal(source)
		json.Unmarshal(jsonString, &event)
		results = append(results, event)
	}

	var responses dto.SearchEventResponse

	// if there does not exist data in the hits, then we can return an empty response
	if len(hits) == 0 {
		responses = dto.SearchEventResponse{
			TotalEvent: 0,
			Events:     []models.EventDocument{},
		}

		return responses, nil
	}

	// if there exist data in the hits, then we can get the total hits
	totalHits := result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)
	responses = dto.SearchEventResponse{
		TotalEvent: int(totalHits),
		Events:     results,
	}

	return responses, nil
}

func buildSearchQuery(query models.SearchQuery) map[string]interface{} {
	startDate, endDate := utils.GetDateRange(query.DateRange)

	// Construct the query map based on the filters
	searchQuery := make(map[string]interface{})
	boolQuery := make(map[string]interface{})
	var must []map[string]interface{}

	if query.Q != "" {
		must = append(must, map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  query.Q,
				"fields": []string{"name", "description", "location"},
				// "type":                 "most_fields", // Can be changed to "best_fields" / "most_fields" / "cross_fields" / "phrase" / "phrase_prefix" for optimization
				"fuzziness":            "AUTO",
				"operator":             "or",
				"fuzzy_transpositions": true,
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
	if query.Category == "all" {
		must = append(must, map[string]interface{}{
			"match_all": map[string]interface{}{},
		})
	} else {
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
	if query.Price != "" {
		must = append(must, map[string]interface{}{
			"match": map[string]interface{}{
				"price": query.Price,
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
	if query.Page != 0 {
		searchQuery["from"] = (query.Page - 1) * query.Offset
	}
	if query.Offset != 0 {
		searchQuery["size"] = query.Offset
	}

	boolQuery["must"] = must
	searchQuery["query"] = map[string]interface{}{
		"bool": boolQuery,
	}

	searchQuery["min_score"] = 0.65 // filter out low score results

	sort := []map[string]interface{}{
		{
			"startDate": map[string]interface{}{
				"order": "desc",
			},
		},
	}
	searchQuery["sort"] = sort

	return searchQuery
}
