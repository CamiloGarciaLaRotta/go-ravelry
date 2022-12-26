package ravelry

import (
	"encoding/json"
	"fmt"

	"github.com/CamiloGarciaLaRotta/go-ravelry/pkg/model"
)

func (client *Client) CurrentUser() (*model.User, error) {
	data, err := client.Api.Get("current_user.json", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get current user: %w", err)
	}

	type currentUserResponse struct {
		User model.User `json:"user"`
	}

	var res currentUserResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal current user: %w", err)
	}

	return &res.User, nil
}
