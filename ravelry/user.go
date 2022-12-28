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
		return nil, model.ErrNoUserID
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

// UpdateUser submits the input user to the server to be updated.
// https://www.ravelry.com/api#people_update
func (client *Client) UpdateUser(user *model.User) (*model.User, error) {
	if user == nil || user.ID == 0 {
		return nil, model.ErrNoUserID
	}

	body, err := json.Marshal(user)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user: %w", err)
	}

	data, err := client.API.Post(fmt.Sprintf("people/%d.json", user.ID), body)
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
