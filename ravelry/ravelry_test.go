package ravelry_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/CamiloGarciaLaRotta/go-ravelry/internal/auth"
	"github.com/CamiloGarciaLaRotta/go-ravelry/internal/testingsupport"
	"github.com/CamiloGarciaLaRotta/go-ravelry/ravelry"
)

func TestNew(t *testing.T) {
	api := ravelry.NewAPI(&testingsupport.FakeAuth{}, "")
	auth := ravelry.NewBasicAuth("foo", "bar")

	ravelry := ravelry.New(api, auth)
	require.NotNil(t, ravelry)
}

func TestNewAPI(t *testing.T) {
	api := ravelry.NewAPI(&testingsupport.FakeAuth{}, "")
	require.NotNil(t, api)
}

func TestNewAuth(t *testing.T) {
	a := ravelry.NewBasicAuth("foo", "bar")
	require.NotNil(t, a)
}

func TestNewAuthFromEnv(t *testing.T) {
	t.Setenv(auth.USER_ENV, "foo")
	t.Setenv(auth.PWD_ENV, "bar")

	a, err := ravelry.NewBasicAuthFromEnv()
	require.NoError(t, err)
	require.NotNil(t, a)
}
