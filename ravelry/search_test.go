package ravelry_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/CamiloGarciaLaRotta/go-ravelry/internal/testingsupport"
	"github.com/CamiloGarciaLaRotta/go-ravelry/pkg/model"
	"github.com/CamiloGarciaLaRotta/go-ravelry/ravelry"
)

func TestSearch_Errors(t *testing.T) {
	fakeAuth := testingsupport.FakeAuth{}
	fakeApi := testingsupport.FakeApi{}
	ravelry := ravelry.New(&fakeApi, &fakeAuth)

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
	fakeAuth := testingsupport.FakeAuth{}
	fakeApi := testingsupport.FakeApi{
		Fail: true,
	}
	ravelry := ravelry.New(&fakeApi, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.Search("foo", 0, nil)
	require.Error(t, err)
	require.Empty(t, res)
}

func TestSearch_UnmarshalError(t *testing.T) {
	fakeAuth := testingsupport.FakeAuth{}
	fakeApi := testingsupport.FakeApi{
		// we return an unexpected empty response
		FakeResp: []byte(""),
	}
	ravelry := ravelry.New(&fakeApi, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.Search("foo", 0, nil)
	require.Error(t, err)
	require.Empty(t, res)
}

func TestSearch(t *testing.T) {
	fakeAuth := testingsupport.FakeAuth{}
	fakeApi := testingsupport.FakeApi{
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
	ravelry := ravelry.New(&fakeApi, &fakeAuth)

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
