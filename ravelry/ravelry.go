package ravelry

import (
	"github.com/CamiloGarciaLaRotta/go-ravelry/internal/api"
	"github.com/CamiloGarciaLaRotta/go-ravelry/internal/auth"
)

// Client is the means by which users can interact with the Ravelry API.
// It takes in a transport layer (Api) and an authenticating layer (Auth).
type Client struct {
	Api  api.API
	Auth auth.Auth
}

// New App that will communicate with the Ravelry API.
func New(api api.API, auth auth.Auth) *Client {
	return &Client{
		Api:  api,
		Auth: auth,
	}
}
