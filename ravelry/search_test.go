package ravelry_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/CamiloGarciaLaRotta/go-ravelry/internal/testingsupport"
	"github.com/CamiloGarciaLaRotta/go-ravelry/pkg/model"
	"github.com/CamiloGarciaLaRotta/go-ravelry/ravelry"
)

func TestSearch_Errors(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// empty query
	res, err := ravelry.Search("", 0, nil)
	require.Error(t, err)
	require.Empty(t, res)

	// negative limit
	res, err = ravelry.Search("foo", -1, nil)
	require.Error(t, err)
	require.Empty(t, res)

	// limit greater than max
	res, err = ravelry.Search("foo", model.SearchLimitMax+1, nil)
	require.Error(t, err)
	require.Empty(t, res)
}

func TestSearch_NetworkError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		Fail: true,
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.Search("foo", 0, nil)
	require.Error(t, err)
	require.Empty(t, res)
}

func TestSearch_UnmarshalError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		// we return an unexpected empty response
		FakeResp: []byte(""),
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.Search("foo", 0, nil)
	require.Error(t, err)
	require.Empty(t, res)
}

func TestSearch(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		FakeResp: []byte(`{
			"results": [
				{
					"title": "Foo",
					"caption": null,
					"type_name": "foo",
					"tiny_image_url": "https://foo.png",
					"image_url": "https://foo.png",
					"record": {
						"type": "PatternSource",
						"id": 78590,
						"permalink": "foo",
						"uri": "/patterns/sources/foo.json"
					}
				},
				{
					"title": "Bar",
					"caption": null,
					"type_name": "bar",
					"tiny_image_url": "https://bar.png",
					"image_url": "https://bar.png",
					"record": {
						"type": "PatternSource",
						"id": 234275,
						"permalink": "bar",
						"uri": "/patterns/sources/bar.json"
					}
				}
			]
		}`),
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	res, err := ravelry.Search("foo", 10, []string{model.SearchTypeShop, model.SearchTypePattern})
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Len(t, res, 2)

	foo := model.SearchObject{
		Title:        "Foo",
		TypeName:     "foo",
		TinyImageURL: "https://foo.png",
		ImageURL:     "https://foo.png",
		Record: model.SearchRecord{
			ID:        78590,
			Type:      "PatternSource",
			Permalink: "foo",
			URI:       "/patterns/sources/foo.json",
		},
	}
	bar := model.SearchObject{
		Title:        "Bar",
		TypeName:     "bar",
		TinyImageURL: "https://bar.png",
		ImageURL:     "https://bar.png",
		Record: model.SearchRecord{
			ID:        234275,
			Type:      "PatternSource",
			Permalink: "bar",
			URI:       "/patterns/sources/bar.json",
		},
	}

	require.Equal(t, res[0], foo)
	require.Equal(t, res[1], bar)
}

func TestSavedSearches_NetworkError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		Fail: true,
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.SavedSearches()
	require.Error(t, err)
	require.Empty(t, res)
}

func TestSavedSearches_UnmarshalError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		// we return an unexpected empty response
		FakeResp: []byte(""),
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.SavedSearches()
	require.Error(t, err)
	require.Empty(t, res)
}

func TestSavedSearches(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		FakeResp: []byte(`{
			"saved_searches": [
				{
					"created_at": "foo",
					"id": 938236,
					"last_loaded": "foo",
					"search_type": "patterns",
					"subscribed": false,
					"updated_at": "foo",
					"title": "crochet",
					"subscription_updated_at": null,
					"search_path": "/patterns/search.json",
					"search_parameters": {
						"foo": "bar"
					}
				},
				{
					"created_at": "bar",
					"id": 938238,
					"last_loaded": "bar",
					"search_type": "patterns",
					"subscribed": true,
					"updated_at": "bar",
					"title": "amigurumi",
					"subscription_updated_at": null,
					"search_path": "/patterns/search.json",
					"search_parameters": {
						"bar": "baz"
					}
				}
			]
		}`),
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	res, err := ravelry.SavedSearches()
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Len(t, res, 2)

	foo := model.SavedSearch{
		CreatedAt:  "foo",
		ID:         938236,
		LastLoaded: "foo",
		SearchType: "patterns",
		UpdatedAt:  "foo",
		Title:      "crochet",
		SearchPath: "/patterns/search.json",
		SearchParams: map[string]string{
			"foo": "bar",
		},
	}
	bar := model.SavedSearch{
		CreatedAt:  "bar",
		ID:         938238,
		Subscribed: true,
		LastLoaded: "bar",
		SearchType: "patterns",
		UpdatedAt:  "bar",
		Title:      "amigurumi",
		SearchPath: "/patterns/search.json",
		SearchParams: map[string]string{
			"bar": "baz",
		},
	}

	require.Equal(t, res[0], foo)
	require.Equal(t, res[1], bar)
}
