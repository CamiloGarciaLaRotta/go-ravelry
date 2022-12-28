package ravelry

import (
	"encoding/json"
	"fmt"

	"github.com/CamiloGarciaLaRotta/go-ravelry/pkg/model"
)

func (client *Client) CurrentUser() (*model.User, error) {
	data, err := client.API.Get("current_user.json", nil)
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

// User fetches the user identified by either the numerical or username passed as id.
// https://www.ravelry.com/api#people_show
func (client *Client) User(id string) (*model.User, error) {
	if id == "" {
		return nil, model.ErrEmptyQuery
	}

	data, err := client.API.Get(fmt.Sprintf("people/%s.json", id), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	type userResponse struct {
		User model.User `json:"user"`
	}

	var res userResponse

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user: %w", err)
	}

	return &res.User, nil
}
