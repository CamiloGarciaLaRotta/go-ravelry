package ravelry

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/CamiloGarciaLaRotta/go-ravelry/pkg/model"
)

func (client *Client) YarnCompanies(query string, page, pageSize int) (*model.YarnCompanySearchResponse, error) {
	queryParams := map[string]string{}

	if strings.TrimSpace(query) != "" {
		queryParams[model.SearchQueryParamQuery] = query
	}

	if page != 0 {
		queryParams[model.SearchQueryParamPage] = fmt.Sprintf("%d", page)
	}

	if pageSize != 0 {
		queryParams[model.SearchQueryParamPageSize] = fmt.Sprintf("%d", pageSize)
	}

	data, err := client.API.Get("yarn_companies/search.json", queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to search for yarn companies: %w", err)
	}

	var res model.YarnCompanySearchResponse

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal yarn company search response: %w", err)
	}

	return &res, nil
}
