package ravelry_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/CamiloGarciaLaRotta/go-ravelry/internal/testingsupport"
	"github.com/CamiloGarciaLaRotta/go-ravelry/pkg/model"
	"github.com/CamiloGarciaLaRotta/go-ravelry/ravelry"
)

func TestCurrentUser_NetworkError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		Fail: true,
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.CurrentUser()
	require.Error(t, err)
	require.Nil(t, res)
}

func TestCurrentUser_UnmarshalError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		// we return an unexpected empty response
		FakeResp: []byte(""),
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.CurrentUser()
	require.Error(t, err)
	require.Nil(t, res)
}

func TestCurrentUser(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
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
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	user := model.User{
		ID:       123,
		Username: "cegal",
	}

	res, err := ravelry.CurrentUser()
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, &user, res)
}
