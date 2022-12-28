package ravelry

import (
	"encoding/json"
	"fmt"

	"github.com/CamiloGarciaLaRotta/go-ravelry/pkg/model"
)

func (client *Client) FiberCategories() ([]model.FiberCategory, error) {
	data, err := client.API.Get("fiber_categories/list.json", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get fiber categories: %w", err)
	}

	type fiberCategoriesResponse struct {
		FiberCategories model.FiberCategory `json:"fiber_categories"`
	}

	var res fiberCategoriesResponse

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal fiber categories: %w", err)
	}

	// this is a bug in the API: the server responds with only 1 fiber category
	// and it's children are all of the actual categories...
	return res.FiberCategories.Children, nil
}
