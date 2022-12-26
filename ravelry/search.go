package ravelry

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/CamiloGarciaLaRotta/go-ravelry/pkg/model"
)

func (client *Client) Search(query string, limit int, types []string) ([]model.SearchObject, error) {
	if strings.TrimSpace(query) == "" {
		return nil, errors.New("search query can't be empty")
	}

	if limit < 0 {
		return nil, errors.New("search limit can't be negative")
	}
	if limit > model.SearchLimitMax {
		return nil, fmt.Errorf("search limit max is %d, got %d", model.SearchLimitMax, limit)
	}
	if limit == 0 {
		limit = model.SearchLimitDefault
	}

	queryParams := map[string]string{
		model.SearchQueryParamQuery: query,
		model.SearchQueryParamLimit: fmt.Sprintf("%d", limit),
	}

	if types != nil {
		queryParams[model.SearchQueryParamType] = strings.Join(types, " ")
	}

	data, err := client.Api.Get("search.json", queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to search: %w", err)
	}

	type searchResponse struct {
		Results []model.SearchObject `json:"results"`
	}

	var res searchResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal search response: %w", err)
	}

	return res.Results, nil
}
