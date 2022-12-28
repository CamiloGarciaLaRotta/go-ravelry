// This package provides the network layer for the Ravelry client.
// It relies on the default http.Client and can be extended to perform additional HTTP methods.
package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/CamiloGarciaLaRotta/go-ravelry/internal/auth"
)

const (
	// API domain.
	RavelryDomain = "https://api.ravelry.com"
	// Timeout for all network requests.
	RequestTimeout = 5 * time.Second
)

var errHTTPStatus = fmt.Errorf("got non %d status", http.StatusOK)

// API defines all the HTTP methods needed to interact with the Ravelry API.
// Defining the interface allows us to mock the network layer in tests.
type API interface {
	Get(url string, params map[string]string) ([]byte, error)
	Post(url string, body []byte) ([]byte, error)
}

// DefaultAPI uses the default http.Client to perform HTTP requests to the Ravelry API.
type DefaultAPI struct {
	auth   auth.Auth
	domain string
}

// New Api which will authenticate to either the default Ravelry API or to an optional alternative domain.
// This alternative domain is useful for tests and local development.
func New(someAuth auth.Auth, alternativeDomain string) *DefaultAPI {
	var domain string
	if alternativeDomain != "" {
		domain = alternativeDomain
	} else {
		domain = RavelryDomain
	}

	api := DefaultAPI{
		auth:   someAuth,
		domain: domain,
	}

	return &api
}

// Get performs a GET request with a default HTTP client and returns the response body.
func (api *DefaultAPI) Get(endpoint string, params map[string]string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", api.domain, endpoint)

	ctx, cancel := context.WithTimeout(context.Background(), RequestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %w", err)
	}

	addQueryParams(req, params)

	api.auth.SetAuth(req)
	req.Header.Set("Content-Type", "application/json")

	client := clientWithTimeout()

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client failed to do request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errHTTPStatus
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return data, nil
}

// Post performs a POST request with a default HTTP client and returns the response body.
func (api *DefaultAPI) Post(endpoint string, body []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", api.domain, endpoint)

	ctx, cancel := context.WithTimeout(context.Background(), RequestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create POST request: %w", err)
	}

	api.auth.SetAuth(req)
	req.Header.Set("Content-Type", "application/json")

	client := clientWithTimeout()

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client failed to do request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errHTTPStatus
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return data, nil
}

func clientWithTimeout() *http.Client {
	netTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: RequestTimeout,
		}).Dial,
		TLSHandshakeTimeout: RequestTimeout,
	}

	return &http.Client{
		Timeout:   RequestTimeout,
		Transport: netTransport,
	}
}

func addQueryParams(req *http.Request, params map[string]string) {
	query := req.URL.Query()

	for k, v := range params {
		query.Add(k, v)
	}

	req.URL.RawQuery = query.Encode()
}
