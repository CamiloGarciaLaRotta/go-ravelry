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

func TestUser_InputError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		Fail: true,
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.User("")
	require.ErrorIs(t, err, model.ErrNoUserID)
	require.Nil(t, res)
}

func TestUser_NetworkError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		Fail: true,
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.User("cegal")
	require.Error(t, err)
	require.Nil(t, res)
}

func TestUser_UnmarshalError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		// we return an unexpected empty response
		FakeResp: []byte(""),
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.User("cegal")
	require.Error(t, err)
	require.Nil(t, res)
}

func TestUser(t *testing.T) {
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

	res, err := ravelry.User("cegal")
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, &user, res)
}

func TestUpdateUser_InputError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		Fail: true,
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.UpdateUser(nil)
	require.ErrorIs(t, err, model.ErrNoUserID)
	require.Nil(t, res)

	res, err = ravelry.UpdateUser(&model.User{})
	require.ErrorIs(t, err, model.ErrNoUserID)
	require.Nil(t, res)
}

func TestUpdateUser_NetworkError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		Fail: true,
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.UpdateUser(&model.User{ID: 1})
	require.Error(t, err)
	require.Nil(t, res)
}

func TestUpdateUser_UnmarshalError(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		// we return an unexpected empty response
		FakeResp: []byte(""),
	}
	ravelry := ravelry.New(&fakeAPI, &fakeAuth)

	// bubbles up the error
	res, err := ravelry.UpdateUser(&model.User{ID: 1})
	require.Error(t, err)
	require.Nil(t, res)
}

func TestUpdateUser(t *testing.T) {
	t.Parallel()

	fakeAuth := testingsupport.FakeAuth{}
	fakeAPI := testingsupport.FakeAPI{
		FakeResp: []byte(`{
			"user": {
				"id": 123,
				"username": "new-cegal",
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

	udpatedUser, err := ravelry.UpdateUser(&user)
	require.NoError(t, err)
	require.NotNil(t, udpatedUser)
	require.NotEqual(t, &user, udpatedUser)
	require.Equal(t, "new-cegal", udpatedUser.Username)
}
