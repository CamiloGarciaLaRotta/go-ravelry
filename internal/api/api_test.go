package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type fakeAuth struct{}

func (auth *fakeAuth) SetAuth(_ *http.Request) {}

func TestRaverlyDomain(t *testing.T) {
	api := New(&fakeAuth{}, "")
	require.Equal(t, "https://api.ravelry.com", api.domain)
}

func TestGet_ServerError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
	}))
	defer ts.Close()

	api := New(&fakeAuth{}, ts.URL)
	res, err := api.Get("foo")
	require.Error(t, err)
	require.Empty(t, res)
}

func TestGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, client")
	}))
	defer ts.Close()

	api := New(&fakeAuth{}, ts.URL)
	res, err := api.Get("foo")
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, []byte("Hello, client"), res)
}
