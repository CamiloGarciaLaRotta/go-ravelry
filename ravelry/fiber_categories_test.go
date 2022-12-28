package ravelry_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/CamiloGarciaLaRotta/go-ravelry/internal/testingsupport"
	"github.com/CamiloGarciaLaRotta/go-ravelry/pkg/model"
	"github.com/CamiloGarciaLaRotta/go-ravelry/ravelry"
)

func TestFiberCategories_NetworkError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		Fail: true,
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.FiberCategories()
	require.Error(t, err)
	require.Empty(t, res)
}

func TestFiberCategories_UnmarshalError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		// we return an unexpected empty response
		FakeResp: []byte(""),
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.FiberCategories()
	require.Error(t, err)
	require.Empty(t, res)
}

func TestFiberCategories(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		FakeResp: []byte(`{
			"fiber_categories": {
				"name": null,
				"permalink": null,
				"short_name": null,
				"children": [
					{
						"id": 174,
						"name": "English",
						"permalink": "english",
						"short_name": "English",
						"children": [
							{
								"id": 123,
								"name": "British",
								"permalink": "british",
								"short_name": "British",
								"children": []
							}
						]
					},
					{
						"id": 175,
						"name": "French",
						"permalink": "french",
						"short_name": "French",
						"children": []
					}
				]
			}
		}`),
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	res, err := ravelry.FiberCategories()
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Len(t, res, 2)

	english := model.FiberCategory{
		ID:        174,
		Name:      "English",
		ShortName: "English",
		Permalink: "english",
		Children: []model.FiberCategory{
			{
				ID:        123,
				Name:      "British",
				Permalink: "british",
				ShortName: "British",
				Children:  []model.FiberCategory{},
			},
		},
	}

	french := model.FiberCategory{
		ID:        175,
		Name:      "French",
		ShortName: "French",
		Permalink: "french",
		Children:  []model.FiberCategory{},
	}

	require.Equal(t, res[0], english)
	require.Equal(t, res[1], french)
}
