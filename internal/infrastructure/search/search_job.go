package search

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/opensearch-project/opensearch-go"
)

func SearchJobs(client *opensearch.Client, query models.SearchJobQuery, page int, offset int) (dto.SearchJobResponse, error) {
	searchQuery := buildSearchJobQuery(query)

	// fmt.Println(searchQuery)
	queryBody, err := json.Marshal(searchQuery)
	if err != nil {
		logs.Error(fmt.Sprintf("failed to marshal search query: %v", err))
		return dto.SearchJobResponse{}, err
	}

	res, err := client.Search(
		client.Search.WithIndex("jobs"),
		client.Search.WithBody(bytes.NewReader(queryBody)),
		client.Search.WithContext(context.Background()),
	)
	if err != nil {
		logs.Error(fmt.Sprintf("failed to execute search on: %v", err))
		// return nil, err
		return dto.SearchJobResponse{}, err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	hits, ok := result["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		logs.Error("No search results found")
		// return nil,
		return dto.SearchJobResponse{}, nil
	}

	var results []models.JobDocument
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		job := models.JobDocument{}
		jsonString, _ := json.Marshal(source)
		json.Unmarshal(jsonString, &job)
		results = append(results, job)
	}

	var responses dto.SearchJobResponse

	// if there does not exist data in the hits, then we can return an empty response
	if len(hits) == 0 {
		responses = dto.SearchJobResponse{
			TotalJob: 0,
			Jobs:     []models.JobDocument{},
		}

		return responses, nil
	}

	// if there exist data in the hits, then we can get the total hits
	totalHits := result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)
	responses = dto.SearchJobResponse{
		TotalJob: int(totalHits),
		Jobs:     results,
	}

	return responses, nil
}

func buildSearchJobQuery(query models.SearchJobQuery) map[string]interface{} {
	searchQuery := make(map[string]interface{})
	boolQuery := make(map[string]interface{})
	var must []map[string]interface{}

	if query.Q != "" {
		must = append(must, map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  query.Q,
				"fields": []string{"title", "description", "location", "organization"},
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
				"location": query.Location,
			},
		})
	}
	if query.Categories == "" {
		must = append(must, map[string]interface{}{
			"match_all": map[string]interface{}{},
		})
	} else {
		categories := strings.Split(query.Categories, ",")
		must = append(must, map[string]interface{}{
			"terms": map[string]interface{}{
				"categories": categories,
			},
		})
	}
	if query.Workplace != "" {
		must = append(must, map[string]interface{}{
			"match": map[string]interface{}{
				"workplace": query.Workplace,
			},
		})
	}
	// Filter by work type (full-time, part-time, etc.)
	if query.WorkType != "" {
		must = append(must, map[string]interface{}{
			"match": map[string]interface{}{
				"workType": query.WorkType,
			},
		})
	}
	// Filter by career stage (entry-level, mid, senior, etc.)
	if query.CareerStage != "" {
		must = append(must, map[string]interface{}{
			"match": map[string]interface{}{
				"careerStage": query.CareerStage,
			},
		})
	}
	if query.SalaryLowerBound != 0 || query.SalaryUpperBound != 0 {
		salaryRange := make(map[string]interface{})
		if query.SalaryLowerBound != 0 {
			salaryRange["gte"] = query.SalaryLowerBound
		}
		if query.SalaryUpperBound != 0 {
			salaryRange["lte"] = query.SalaryUpperBound
		}
		must = append(must, map[string]interface{}{
			"range": map[string]interface{}{
				"salary": salaryRange,
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

	return searchQuery
}
