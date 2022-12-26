package testingsupport

import (
	"errors"
	"net/http"
)

// FakeAuth is a test mock for the auth.Auth interface.
type FakeAuth struct{}

// SetAuth is a mock implementation for tests.
func (auth *FakeAuth) SetAuth(_ *http.Request) {}

// FakeApi is a test mock for the api.API interface.
type FakeApi struct {
	// whether the call to the API layer should fail or not
	Fail bool
	// the stubbed out server response
	FakeResp []byte
}

// Get is a mock implementation for tests.
func (api *FakeApi) Get(url string) ([]byte, error) {
	if api.Fail {
		return nil, errors.New("booom")
	}
	return api.FakeResp, nil
}
