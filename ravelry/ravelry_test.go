package ravelry_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/CamiloGarciaLaRotta/go-ravelry/internal/auth"
	"github.com/CamiloGarciaLaRotta/go-ravelry/internal/testingsupport"
	"github.com/CamiloGarciaLaRotta/go-ravelry/pkg/model"
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

func TestNewBasicAuth(t *testing.T) {
	a := ravelry.NewBasicAuth("foo", "bar")
	require.NotNil(t, a)
}

func TestNewBasicAuthFromEnv(t *testing.T) {
	t.Setenv(auth.USER_ENV, "foo")
	t.Setenv(auth.PWD_ENV, "bar")

	a, err := ravelry.NewBasicAuthFromEnv()
	require.NoError(t, err)
	require.NotNil(t, a)
}

func TestReadOnlyEndpoint(t *testing.T) {
	// we expect the ENV vars to be present in localhost and CI
	auth, err := ravelry.NewBasicAuthFromEnv()
	require.NoError(t, err)

	api := ravelry.NewAPI(auth, "")
	ravelry := ravelry.New(api, auth)

	colors, err := ravelry.ColorFamilies()
	require.NoError(t, err)
	require.NotEmpty(t, colors)
}

func TestPersonalEndpoint(t *testing.T) {
	// we expect the ENV vars to be present in localhost and CI
	auth, err := ravelry.NewBasicAuthFromEnv()
	require.NoError(t, err)

	api := ravelry.NewAPI(auth, "")
	ravelry := ravelry.New(api, auth)

	colors, err := ravelry.CurrentUser()
	require.NoError(t, err)
	require.NotEmpty(t, colors)
}

func TestURLParamEndpoint(t *testing.T) {
	// we expect the ENV vars to be present in localhost and CI
	auth, err := ravelry.NewBasicAuthFromEnv()
	require.NoError(t, err)

	api := ravelry.NewAPI(auth, "")
	ravelry := ravelry.New(api, auth)

	colors, err := ravelry.Search("foo", 1, []string{model.SearchTypeShop, model.SearchTypePatternAuthor})
	require.NoError(t, err)
	require.NotEmpty(t, colors)
}

func TestYarnCompaniesEndpoint(t *testing.T) {
	// we expect the ENV vars to be present in localhost and CI
	auth, err := ravelry.NewBasicAuthFromEnv()
	require.NoError(t, err)

	api := ravelry.NewAPI(auth, "")
	ravelry := ravelry.New(api, auth)

	colors, err := ravelry.YarnCompanies("puppy", 1, 2)
	require.NoError(t, err)
	require.NotEmpty(t, colors.Companies)
}

func TestYarnAttributesEndpoint(t *testing.T) {
	// we expect the ENV vars to be present in localhost and CI
	auth, err := ravelry.NewBasicAuthFromEnv()
	require.NoError(t, err)

	api := ravelry.NewAPI(auth, "")
	ravelry := ravelry.New(api, auth)

	attrs, err := ravelry.YarnAttributes()
	require.NoError(t, err)
	require.NotEmpty(t, attrs)
}
