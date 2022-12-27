package ravelry

import (
	"encoding/json"
	"fmt"

	"github.com/CamiloGarciaLaRotta/go-ravelry/pkg/model"
)

func (client *Client) ColorFamilies() ([]model.ColorFamily, error) {
	data, err := client.API.Get("color_families.json", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get color families: %w", err)
	}

	type colorFamilyResponse struct {
		ColorFamilies []model.ColorFamily `json:"color_families"`
	}

	var res colorFamilyResponse

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal color families: %w", err)
	}

	return res.ColorFamilies, nil
}
