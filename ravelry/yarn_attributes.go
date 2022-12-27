package ravelry

import (
	"encoding/json"
	"fmt"

	"github.com/CamiloGarciaLaRotta/go-ravelry/pkg/model"
)

func (client *Client) YarnAttributes() ([]model.YarnAttributeParent, error) {
	data, err := client.Api.Get("yarn_attributes/groups.json", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get yarn attributes: %w", err)
	}

	type yarnAttributeResponse struct {
		Attrs []model.YarnAttributeParent `json:"yarn_attribute_groups"`
	}

	var res yarnAttributeResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal yarn attribute response: %w", err)
	}

	return res.Attrs, nil
}
