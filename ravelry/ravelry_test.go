package ravelry_test

import (
	"testing"

	"github.com/CamiloGarciaLaRotta/go-ravelry/internal/auth"
	"github.com/CamiloGarciaLaRotta/go-ravelry/ravelry"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	api := ravelry.NewAPI(&fakeAuth{}, "")
	auth := ravelry.NewAuth("foo", "bar")

	ravelry := ravelry.New(api, auth)
	require.NotNil(t, ravelry)
}

func TestNewAPI(t *testing.T) {
	api := ravelry.NewAPI(&fakeAuth{}, "")
	require.NotNil(t, api)
}

func TestNewAuth(t *testing.T) {
	a := ravelry.NewAuth("foo", "bar")
	require.NotNil(t, a)
}

func TestNewAuthFromEnv(t *testing.T) {
	t.Setenv(auth.USER_ENV, "foo")
	t.Setenv(auth.PWD_ENV, "bar")

	a, err := ravelry.NewAuthFromEnv()
	require.NoError(t, err)
	require.NotNil(t, a)
}
