package ravelry

import (
	"encoding/json"
	"fmt"

	"github.com/CamiloGarciaLaRotta/go-ravelry/pkg/model"
)

func (client *Client) YarnWeights() ([]model.YarnWeight, error) {
	data, err := client.Api.Get("yarn_weights.json", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get yarn weights: %w", err)
	}

	type yarnWeightResponse struct {
		YarnWeights []model.YarnWeight `json:"yarn_weights"`
	}

	var res yarnWeightResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal yarn weights: %w", err)
	}

	return res.YarnWeights, nil
}
