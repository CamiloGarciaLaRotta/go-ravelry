package ravelry_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/CamiloGarciaLaRotta/go-ravelry/internal/testingsupport"
	"github.com/CamiloGarciaLaRotta/go-ravelry/pkg/model"
	"github.com/CamiloGarciaLaRotta/go-ravelry/ravelry"
)

func TestYarnAttributes_NetworkError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		Fail: true,
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.YarnAttributes()
	require.Error(t, err)
	require.Empty(t, res)
}

func TestYarnAttributes_UnmarshalError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		// we return an unexpected empty response
		FakeResp: []byte(""),
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.YarnAttributes()
	require.Error(t, err)
	require.Empty(t, res)
}

func TestYarnAttributes(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		FakeResp: []byte(`{
			"yarn_attribute_groups": [
				{
					"id": 9,
					"name": "foo",
					"permalink": "foo",
					"yarn_attributes": [
						{
							"created_at": "foo_attr",
							"description": "foo_attr",
							"id": 7,
							"name": "foo_attr",
							"permalink": "foo_attr",
							"sort_order": 0,
							"yarn_attribute_group_id": 9
						}
					],
					"children": []
				},
				{
					"id": 13,
					"name": "bar",
					"permalink": "bar",
					"yarn_attributes": [],
					"children": [
						{
							"id": 14,
							"name": "bar_child",
							"permalink": "bar_child",
							"yarn_attributes": [
								{
								"created_at": "bar_child_attr",
								"description": "bar_child_attr",
								"id": 30,
								"name": "bar_child_attr",
								"permalink": "bar_child_attr",
								"sort_order": 0,
								"yarn_attribute_group_id": 14
								}
							],
							"children": []
						}
					]
				}
			]
		}`),
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	res, err := ravelry.YarnAttributes()
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Len(t, res, 2)

	foo := model.YarnAttributeParent{
		ID:        9,
		Name:      "foo",
		Permalink: "foo",
		YarnAttributes: []model.YarnAttributeNode{
			{
				ID:                   7,
				SortOrder:            0,
				YarnAttributeGroupID: 9,
				CreatedAt:            "foo_attr",
				Description:          "foo_attr",
				Name:                 "foo_attr",
				Permalink:            "foo_attr",
			},
		},
		Children: []model.YarnAttributeParent{},
	}
	bar := model.YarnAttributeParent{
		ID:             13,
		Name:           "bar",
		Permalink:      "bar",
		YarnAttributes: []model.YarnAttributeNode{},
		Children: []model.YarnAttributeParent{
			{
				ID:        14,
				Name:      "bar_child",
				Permalink: "bar_child",
				YarnAttributes: []model.YarnAttributeNode{
					{
						ID:                   30,
						SortOrder:            0,
						YarnAttributeGroupID: 14,
						CreatedAt:            "bar_child_attr",
						Description:          "bar_child_attr",
						Name:                 "bar_child_attr",
						Permalink:            "bar_child_attr",
					},
				},
				Children: []model.YarnAttributeParent{},
			},
		},
	}

	require.Equal(t, res[0], foo)
	require.Equal(t, res[1], bar)
}
