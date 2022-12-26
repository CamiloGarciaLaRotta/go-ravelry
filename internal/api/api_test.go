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
	res, err := api.Get("foo", nil)
	require.Error(t, err)
	require.Empty(t, res)
}

func TestGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, client")
	}))
	defer ts.Close()

	api := New(&fakeAuth{}, ts.URL)
	res, err := api.Get("foo", nil)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, []byte("Hello, client"), res)
}

func TestGet_Params(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		require.Equal(t, "bar", q.Get("foo"))
	}))
	defer ts.Close()

	api := New(&fakeAuth{}, ts.URL)
	_, err := api.Get("foo", map[string]string{"foo": "bar"})
	require.NoError(t, err)
}
