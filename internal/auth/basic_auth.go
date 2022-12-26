// This package defines the means by which a user can [authenticate] with the Ravelry API.
//
// Ravelry supports Basic-Auth and Oauth.
// We currently only support Basic-Auth (read-only and personal).
//
// [authenticate]: https://www.ravelry.com/api#authenticating
package auth

import (
	"errors"
	"net/http"
	"os"
)

const (
	// Username env var
	USER_ENV = "RAVELRY_USER"
	// Password env var
	PWD_ENV = "RAVELRY_PWD"
)

// Auth defines all the means by which a user can authenticate with the Ravelry API.
// Defining the interface allows us to mock the network layer in tests.
type Auth interface {

	// set the headers required to auth with the API
	SetAuth(req *http.Request)
}

// BasicAuth auth method as described in
// https://www.ravelry.com/api#authenticating.
// It can be used to access read-only and personal endpoints.
type BasicAuth struct {
	user, pass string
}

// SetAuth injects the headers required for basic auth
func (auth *BasicAuth) SetAuth(req *http.Request) {
	req.SetBasicAuth(auth.user, auth.pass)
}

// NewBasicAuth takes directly the user and password
func NewBasicAuth(u, p string) *BasicAuth {
	return &BasicAuth{user: u, pass: p}
}

// NewBasicAuthFromEnv extracts the user and password from ENV:
//   - user: $RAVELRY_USER.
//   - pass: $RAVELRY_PWD.
func NewBasicAuthFromEnv() (*BasicAuth, error) {
	u := os.Getenv(USER_ENV)
	if u == "" {
		return nil, errors.New("$RAVELRY_USER is not defined")
	}
	p := os.Getenv(PWD_ENV)
	if p == "" {
		return nil, errors.New("$RAVELRY_PWD is not defined")
	}

	return &BasicAuth{user: u, pass: p}, nil
}
