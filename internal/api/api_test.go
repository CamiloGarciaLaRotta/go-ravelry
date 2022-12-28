package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type fakeAuth struct{}

func (auth *fakeAuth) SetAuth(_ *http.Request) {}

func TestRaverlyDomain(t *testing.T) {
	t.Parallel()

	api := New(&fakeAuth{}, "")
	require.Equal(t, "https://api.ravelry.com", api.domain)
}

func TestGet_ServerError(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	api := New(&fakeAuth{}, ts.URL)
	res, err := api.Get("foo", nil)
	require.ErrorIs(t, err, errHTTPStatus)
	require.Empty(t, res)
}

func TestGet(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		require.Equal(t, "bar", q.Get("foo"))
	}))
	defer server.Close()

	api := New(&fakeAuth{}, server.URL)
	_, err := api.Get("foo", map[string]string{"foo": "bar"})
	require.NoError(t, err)
}

func TestPost_ServerError(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	api := New(&fakeAuth{}, ts.URL)
	res, err := api.Post("foo", nil)
	require.ErrorIs(t, err, errHTTPStatus)
	require.Empty(t, res)
}

func TestPost(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, client")
	}))
	defer ts.Close()

	api := New(&fakeAuth{}, ts.URL)
	res, err := api.Post("foo", nil)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, []byte("Hello, client"), res)
}

func TestPost_Params(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		log.Println(string(body))
		require.Equal(t, []byte("foobar"), body)
	}))
	defer server.Close()

	api := New(&fakeAuth{}, server.URL)
	_, err := api.Post("foo", []byte("foobar"))
	require.NoError(t, err)
}
