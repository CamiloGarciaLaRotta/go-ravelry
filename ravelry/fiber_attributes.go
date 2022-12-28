package ravelry

import (
	"encoding/json"
	"fmt"

	"github.com/CamiloGarciaLaRotta/go-ravelry/pkg/model"
)

func (client *Client) FiberAttributes() ([]model.FiberAttribute, error) {
	data, err := client.API.Get("fiber_attributes.json", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get fiber attributes: %w", err)
	}

	type fiberAttrsResponse struct {
		FiberAttrs []model.FiberAttribute `json:"fiber_attributes"`
	}

	var res fiberAttrsResponse

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal fiber attributes: %w", err)
	}

	return res.FiberAttrs, nil
}

func (client *Client) FiberAttributeGroups() ([]model.FiberAttributeGroup, error) {
	data, err := client.API.Get("fiber_attribute_groups.json", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get fiber attribute categories: %w", err)
	}

	type fiberAttrGroupResponse struct {
		Groups []model.FiberAttributeGroup `json:"fiber_attribute_groups"`
	}

	var res fiberAttrGroupResponse

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal fiber attribute categories: %w", err)
	}

	return res.Groups, nil
}
