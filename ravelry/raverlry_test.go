package ravelry_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/CamiloGarciaLaRotta/go-ravelry/ravelry"
)

func TestReadOnlyEndpoint(t *testing.T) {
	auth, err := ravelry.NewBasicAuthFromEnv()
	require.NoError(t, err)

	api := ravelry.NewAPI(auth, "")
	ravelry := ravelry.New(api, auth)

	colors, err := ravelry.ColorFamilies()
	require.NoError(t, err)
	require.NotEmpty(t, colors)
}

func TestPersonalEndpoint(t *testing.T) {
	auth, err := ravelry.NewBasicAuthFromEnv()
	require.NoError(t, err)

	api := ravelry.NewAPI(auth, "")
	ravelry := ravelry.New(api, auth)

	colors, err := ravelry.CurrentUser()
	require.NoError(t, err)
	require.NotEmpty(t, colors)
}
