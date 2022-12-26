package auth_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/CamiloGarciaLaRotta/go-ravelry/internal/auth"
)

func TestNewFromEnv_Errors(t *testing.T) {
	t.Setenv(auth.USER_ENV, "")
	t.Setenv(auth.PWD_ENV, "")

	a, err := auth.NewFromEnv()
	require.Error(t, err)
	require.Nil(t, a)

	t.Setenv(auth.USER_ENV, "foo")

	a, err = auth.NewFromEnv()
	require.Error(t, err)
	require.Nil(t, a)
}

func TestNewFromEnv(t *testing.T) {
	t.Setenv(auth.USER_ENV, "foo")
	t.Setenv(auth.PWD_ENV, "bar")

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	require.NoError(t, err)

	a, err := auth.NewFromEnv()
	require.NoError(t, err)

	a.SetAuth(r)

	u, p, ok := r.BasicAuth()
	require.True(t, ok)
	require.Equal(t, "foo", u)
	require.Equal(t, "bar", p)
}

func TestNew(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	require.NoError(t, err)

	a := auth.New("foo", "bar")

	a.SetAuth(r)

	u, p, ok := r.BasicAuth()
	require.True(t, ok)
	require.Equal(t, "foo", u)
	require.Equal(t, "bar", p)
}
