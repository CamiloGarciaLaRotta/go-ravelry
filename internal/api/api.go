// This package provides the network layer for the Ravelry client.
// It relies on the default http.Client and can be extended to perform additional HTTP methods.
package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/CamiloGarciaLaRotta/go-ravelry/internal/auth"
)

const RAVELRY_DOMAIN = "https://api.ravelry.com"

// API defines all the HTTP methods needed to interact with the Ravelry API.
// Defining the interface allows us to mock the network layer in tests.
type API interface {
	Get(url string, params map[string]string) ([]byte, error)
}

type Api struct {
	auth   auth.Auth
	domain string
}

// New Api which will authenticate to either the default Ravelry API or to an optional alternative domain.
// This alternative domain is useful for tests and local development.
func New(a auth.Auth, alternativeDomain string) *Api {
	var d string
	if alternativeDomain != "" {
		d = alternativeDomain
	} else {
		d = RAVELRY_DOMAIN
	}
	return &Api{auth: a, domain: d}
}

// Get performs a GET request with a default HTTP client and returns the response body.
func (api *Api) Get(endpoint string, params map[string]string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", api.domain, endpoint)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %w", err)
	}

	addQueryParams(req, params)

	api.auth.SetAuth(req)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s got HTTP status %d", endpoint, res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return data, nil
}

func addQueryParams(r *http.Request, params map[string]string) {
	q := r.URL.Query()

	for k, v := range params {
		q.Add(k, v)
	}

	r.URL.RawQuery = q.Encode()
}
