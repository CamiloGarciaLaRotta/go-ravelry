// This package provides the constructors needed to setup the [Ravelry API] client.
// It allows instantiation of the network layer (api) and the authentication layer (auth).
// Both of which are required for the client to be spun up.
//
// [Ravelry API]: https://www.ravelry.com/api
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

// NewAPI creates the network layer for the Ravelry client.
// It will authenticate to either the default Ravelry API or to an optional alternative domain.
// This alternative domain is useful for tests and local development.
func NewAPI(a auth.Auth, alternativeDomain string) *api.Api {
	return api.New(a, alternativeDomain)
}

// NewBasicAuth creates the auth layer for the Ravelry client.
// It takes directly the user and password.
func NewBasicAuth(u, p string) *auth.BasicAuth {
	return auth.NewBasicAuth(u, p)
}

// NewBasicAuthFromEnv creates the auth layer for the Ravelry client.
// It extracts the user and password from ENV:
//   - user: $RAVELRY_USER.
//   - pass: $RAVELRY_PWD.
func NewBasicAuthFromEnv() (*auth.BasicAuth, error) {
	return auth.NewBasicAuthFromEnv()
}
