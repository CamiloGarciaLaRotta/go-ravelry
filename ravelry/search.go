package ravelry

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/CamiloGarciaLaRotta/go-ravelry/pkg/model"
)

func (client *Client) Search(query string, limit int, types []string) ([]model.SearchObject, error) {
	if strings.TrimSpace(query) == "" {
		return nil, model.ErrEmptyQuery
	}

	if limit < 0 {
		return nil, model.ErrNegativeLimit
	}

	if limit > model.SearchLimitMax {
		return nil, model.ErrAboveLimitMax
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

	data, err := client.API.Get("search.json", queryParams)
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

func (client *Client) SavedSearches() ([]model.SavedSearch, error) {
	data, err := client.API.Get("saved_searches/list.json", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get saved searches: %w", err)
	}

	type searchResponse struct {
		Searches []model.SavedSearch `json:"saved_searches"`
	}

	var res searchResponse

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal saved searches response: %w", err)
	}

	return res.Searches, nil
}
