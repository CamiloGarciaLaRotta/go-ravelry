package ravelry_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/CamiloGarciaLaRotta/go-ravelry/pkg/model"
	"github.com/CamiloGarciaLaRotta/go-ravelry/ravelry"
)

type fakeAuth struct{}

func (auth *fakeAuth) SetAuth(_ *http.Request) {}

type fakeApi struct {
	fail     bool
	fakeResp []byte
}

func (api *fakeApi) Get(url string) ([]byte, error) {
	if api.fail {
		return nil, errors.New("booom")
	}
	return api.fakeResp, nil
}

func TestColorFamilies_NetworkError(t *testing.T) {
	fakeAuth := fakeAuth{}
	fakeApi := fakeApi{
		fail: true,
	}
	ravelry := ravelry.New(&fakeApi, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.ColorFamilies()
	require.Error(t, err)
	require.Empty(t, res)
}

func TestColorFamilies_UnmarshalError(t *testing.T) {
	fakeAuth := fakeAuth{}
	fakeApi := fakeApi{
		// we return an unexpected empty response
		fakeResp: []byte(""),
	}
	ravelry := ravelry.New(&fakeApi, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.ColorFamilies()
	require.Error(t, err)
	require.Empty(t, res)
}

func TestColorFamilies(t *testing.T) {
	fakeAuth := fakeAuth{}
	fakeApi := fakeApi{
		// we return an unexpected empty response
		fakeResp: []byte(`{
			"color_families": [
				{
					"color": null,
					"id": 1,
					"name": "Yellow",
					"permalink": "Yellow",
					"spectrum_order": 1
				},
				{
					"color": null,
					"id": 3,
					"name": "Orange",
					"permalink": "Orange",
					"spectrum_order": 3
				}
			]
		}`),
	}
	ravelry := ravelry.New(&fakeApi, &fakeAuth)

	res, err := ravelry.ColorFamilies()
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Len(t, res, 2)

	yellow := model.ColorFamily{
		ID:            1,
		Name:          "Yellow",
		Permalink:     "Yellow",
		SpectrumOrder: 1,
	}
	orange := model.ColorFamily{
		ID:            3,
		Name:          "Orange",
		Permalink:     "Orange",
		SpectrumOrder: 3,
	}

	require.Equal(t, res[0], yellow)
	require.Equal(t, res[1], orange)
}
