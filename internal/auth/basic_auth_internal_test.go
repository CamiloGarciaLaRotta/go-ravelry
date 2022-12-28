package auth

import (
	"testing"

	"github.com/stretchr/testify/require"
)

//nolint:paralleltest
func TestNewBasicAuthFromEnv_Errors(t *testing.T) {
	t.Setenv(UserENV, "")
	t.Setenv(KeyENV, "")

	someAuth, err := NewBasicAuthFromEnv()
	require.ErrorIs(t, err, errNoUser)
	require.Nil(t, someAuth)

	t.Setenv(UserENV, "foo")

	someAuth, err = NewBasicAuthFromEnv()
	require.ErrorIs(t, err, errNoKey)
	require.Nil(t, someAuth)
}
