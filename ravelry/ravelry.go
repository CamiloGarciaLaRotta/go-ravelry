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

// New creates the auth layer for the Ravelry client.
// It takes directly the user and password.
func NewAuth(u, p string) *auth.BasicAuth {
	return auth.New(u, p)
}

// NewFromEnv creates the auth layer for the Ravelry client.
// It extracts the user and password from ENV:
// user: $RAVELRY_USER
// pass: $RAVELRY_PWD
func NewAuthFromEnv() (*auth.BasicAuth, error) {
	return auth.NewFromEnv()
}

// NewAPI creates the network layer for the Ravelry client.
// It will authenticate to either the default Ravelry API or to an optional alternative domain.
// This alternative domain is useful for tests and local development.
func NewAPI(a auth.Auth, alternativeDomain string) *api.Api {
	return api.New(a, alternativeDomain)
}
