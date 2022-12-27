package ravelry_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/CamiloGarciaLaRotta/go-ravelry/internal/testingsupport"
	"github.com/CamiloGarciaLaRotta/go-ravelry/pkg/model"
	"github.com/CamiloGarciaLaRotta/go-ravelry/ravelry"
)

func TestYarnCompanies_NetworkError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		Fail: true,
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.YarnCompanies("foo", 0, 0)
	require.Error(t, err)
	require.Empty(t, res)
}

func TestYarnCompanies_UnmarshalError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		// we return an unexpected empty response
		FakeResp: []byte(""),
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.YarnCompanies("foo", 0, 0)
	require.Error(t, err)
	require.Empty(t, res)
}

func TestYarnCompanies(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		FakeResp: []byte(`{"yarn_companies": [
				{
					"id": 3403,
					"name": "Puppy",
					"permalink": "puppy",
					"url": "http://www.puppyarn.com/",
					"yarns_count": 254
				},
				{
					"id": 3691,
					"name": "Roberta di Camerino",
					"permalink": "roberta-di-camerino",
					"url": "http://www.robertadicamerino.com/",
					"yarns_count": 21
				}
			],
			"paginator": {
				"page_count": 1,
				"page": 1,
				"page_size": 48,
				"results": 2,
				"last_page": 1
			}
		}`),
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	res, err := ravelry.YarnCompanies("foo", 1, 2)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Len(t, res.Companies, 2)

	puppy := model.YarnCompany{
		ID:         3403,
		YarnsCount: 254,
		Name:       "Puppy",
		Permalink:  "puppy",
		URL:        "http://www.puppyarn.com/",
	}
	roberta := model.YarnCompany{
		ID:         3691,
		YarnsCount: 21,
		Name:       "Roberta di Camerino",
		Permalink:  "roberta-di-camerino",
		URL:        "http://www.robertadicamerino.com/",
	}

	require.Equal(t, res.Companies[0], puppy)
	require.Equal(t, res.Companies[1], roberta)

	paginator := model.Paginator{
		PageCount: 1,
		Page:      1,
		PageSize:  48,
		Results:   2,
		LastPage:  1,
	}
	require.Equal(t, res.Paginator, paginator)
}
