package ravelry_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/CamiloGarciaLaRotta/go-ravelry/internal/testingsupport"
	"github.com/CamiloGarciaLaRotta/go-ravelry/pkg/model"
	"github.com/CamiloGarciaLaRotta/go-ravelry/ravelry"
)

func TestFiberAttributes_NetworkError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		Fail: true,
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.FiberAttributes()
	require.Error(t, err)
	require.Empty(t, res)
}

func TestFiberAttributes_UnmarshalError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		// we return an unexpected empty response
		FakeResp: []byte(""),
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.FiberAttributes()
	require.Error(t, err)
	require.Empty(t, res)
}

func TestFiberAttributes(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		FakeResp: []byte(`{
			"fiber_attributes": [
				{
					"fiber_attribute_group_id": 1,
					"id": 1,
					"name": "Commercially dyed",
					"permalink": "commercially-dyed"
				},
				{
					"fiber_attribute_group_id": 1,
					"id": 2,
					"name": "Handdyed",
					"permalink": "handdyed"
				}
			]
		}`),
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	res, err := ravelry.FiberAttributes()
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Len(t, res, 2)

	foo := model.FiberAttribute{
		FiberAttrGroupID: 1,
		ID:               1,
		Name:             "Commercially dyed",
		Permalink:        "commercially-dyed",
	}
	bar := model.FiberAttribute{
		FiberAttrGroupID: 1,
		ID:               2,
		Name:             "Handdyed",
		Permalink:        "handdyed",
	}

	require.Equal(t, res[0], foo)
	require.Equal(t, res[1], bar)
}

func TestFiberAttributeCategory_NetworkError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		Fail: true,
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.FiberAttributeGroups()
	require.Error(t, err)
	require.Empty(t, res)
}

func TestFiberAttributeCategory_UnmarshalError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		// we return an unexpected empty response
		FakeResp: []byte(""),
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.FiberAttributeGroups()
	require.Error(t, err)
	require.Empty(t, res)
}

func TestFiberAttributeCategory(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		FakeResp: []byte(`{
			"fiber_attribute_groups": [
				{
					"id": 1,
					"name": "Color",
					"parent_id": null,
					"permalink": "color"
				},
				{
					"id": 2,
					"name": "Fiber Content",
					"parent_id": 1,
					"permalink": "content"
				}
			]
		}`),
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	res, err := ravelry.FiberAttributeGroups()
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Len(t, res, 2)

	color := model.FiberAttributeGroup{
		ID:        1,
		Name:      "Color",
		Permalink: "color",
	}
	fiber := model.FiberAttributeGroup{
		ParentID:  1,
		ID:        2,
		Name:      "Fiber Content",
		Permalink: "content",
	}

	require.Equal(t, res[0], color)
	require.Equal(t, res[1], fiber)
}
