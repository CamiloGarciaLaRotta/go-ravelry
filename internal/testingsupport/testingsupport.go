// This package provides mock implmentations for the auth and api layers.
package testingsupport

import (
	"errors"
	"net/http"
)

// FakeAuth is a test mock for the auth.Auth interface.
type FakeAuth struct{}

// SetAuth is a mock implementation for tests.
func (auth *FakeAuth) SetAuth(_ *http.Request) {}

// FakeAPI is a test mock for the api.API interface.
type FakeAPI struct {
	// whether the call to the API layer should fail or not
	Fail bool
	// the stubbed out server response
	FakeResp []byte
}

var errDummy = errors.New("booom")

// Get is a mock implementation for tests.
func (api *FakeAPI) Get(url string, params map[string]string) ([]byte, error) {
	if api.Fail {
		return nil, errDummy
	}

	return api.FakeResp, nil
}

// Post is a mock implementation for tests.
func (api *FakeAPI) Post(url string, body []byte) ([]byte, error) {
	if api.Fail {
		return nil, errDummy
	}

	return api.FakeResp, nil
}
