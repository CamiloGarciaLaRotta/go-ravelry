package ravelry_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/CamiloGarciaLaRotta/go-ravelry/internal/testingsupport"
	"github.com/CamiloGarciaLaRotta/go-ravelry/pkg/model"
	"github.com/CamiloGarciaLaRotta/go-ravelry/ravelry"
)

func TestYarnWeights_NetworkError(t *testing.T) {
	fakeAuth := testingsupport.FakeAuth{}
	fakeApi := testingsupport.FakeApi{
		Fail: true,
	}
	ravelry := ravelry.New(&fakeApi, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.YarnWeights()
	require.Error(t, err)
	require.Empty(t, res)
}

func TestYarnWeights_UnmarshalError(t *testing.T) {
	fakeAuth := testingsupport.FakeAuth{}
	fakeApi := testingsupport.FakeApi{
		// we return an unexpected empty response
		FakeResp: []byte(""),
	}
	ravelry := ravelry.New(&fakeApi, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.YarnWeights()
	require.Error(t, err)
	require.Empty(t, res)
}

func TestYarnWeights(t *testing.T) {
	fakeAuth := testingsupport.FakeAuth{}
	fakeApi := testingsupport.FakeApi{
		FakeResp: []byte(`{
			"yarn_weights": [
				{
					"crochet_gauge": "",
					"id": 1,
					"knit_gauge": "18",
					"max_gauge": null,
					"min_gauge": null,
					"name": "Aran",
					"ply": "10",
					"wpi": "8"
				},
				{
					"crochet_gauge": "",
					"id": 4,
					"knit_gauge": "14-15",
					"max_gauge": null,
					"min_gauge": null,
					"name": "Bulky",
					"ply": "12",
					"wpi": "7"
				}
			]
		}`),
	}
	ravelry := ravelry.New(&fakeApi, &fakeAuth)

	res, err := ravelry.YarnWeights()
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Len(t, res, 2)

	aran := model.YarnWeight{
		ID:        1,
		Name:      "Aran",
		KnitGauge: "18",
		Ply:       "10",
		Wpi:       "8",
	}
	bulky := model.YarnWeight{
		ID:        4,
		Name:      "Bulky",
		KnitGauge: "14-15",
		Ply:       "12",
		Wpi:       "7",
	}

	require.Equal(t, res[0], aran)
	require.Equal(t, res[1], bulky)
}
