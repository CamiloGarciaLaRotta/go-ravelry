package auth_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/CamiloGarciaLaRotta/go-ravelry/internal/auth"
)

//nolint:paralleltest
func TestNewBasicAuthFromEnv(t *testing.T) {
	t.Setenv(auth.UserENV, "foo")
	t.Setenv(auth.KeyENV, "bar")

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	require.NoError(t, err)

	a, err := auth.NewBasicAuthFromEnv()
	require.NoError(t, err)

	a.SetAuth(req)

	u, p, ok := req.BasicAuth()
	require.True(t, ok)
	require.Equal(t, "foo", u)
	require.Equal(t, "bar", p)
}

func TestNewBasicAuth(t *testing.T) {
	t.Parallel()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	require.NoError(t, err)

	a := auth.NewBasicAuth("foo", "bar")

	a.SetAuth(req)

	u, p, ok := req.BasicAuth()
	require.True(t, ok)
	require.Equal(t, "foo", u)
	require.Equal(t, "bar", p)
}
