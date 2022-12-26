package ravelry_test

import (
	"testing"

	"github.com/CamiloGarciaLaRotta/go-ravelry/internal/testingsupport"
	"github.com/CamiloGarciaLaRotta/go-ravelry/pkg/model"
	"github.com/CamiloGarciaLaRotta/go-ravelry/ravelry"
	"github.com/stretchr/testify/require"
)

func TestCurrentUser_NetworkError(t *testing.T) {
	fakeAuth := testingsupport.FakeAuth{}
	fakeApi := testingsupport.FakeApi{
		Fail: true,
	}
	ravelry := ravelry.New(&fakeApi, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.CurrentUser()
	require.Error(t, err)
	require.Nil(t, res)
}

func TestCurrentUser_UnmarshalError(t *testing.T) {
	fakeAuth := testingsupport.FakeAuth{}
	fakeApi := testingsupport.FakeApi{
		// we return an unexpected empty response
		FakeResp: []byte(""),
	}
	ravelry := ravelry.New(&fakeApi, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.CurrentUser()
	require.Error(t, err)
	require.Nil(t, res)
}

func TestCurrentUser(t *testing.T) {
	fakeAuth := testingsupport.FakeAuth{}
	fakeApi := testingsupport.FakeApi{
		FakeResp: []byte(`{
			"user": {
				"id": 123,
				"username": "cegal",
				"tiny_photo_url": null,
				"small_photo_url": null,
				"photo_url": null,
				"large_photo_url": null
			}
		  }`),
	}
	ravelry := ravelry.New(&fakeApi, &fakeAuth)

	u := model.User{
		ID:       123,
		Username: "cegal",
	}

	res, err := ravelry.CurrentUser()
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, &u, res)
}
