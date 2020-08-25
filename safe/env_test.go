package safe_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "go.octolab.org/safe"
)

func TestSetEnvs(t *testing.T) {
	t.Run("bad call", func(t *testing.T) {
		release, err := SetEnvs()
		assert.Nil(t, release)
		assert.Equal(t, ErrBadCall, err)

		release, err = SetEnvs("a", "b", "c")
		assert.Nil(t, release)
		assert.Equal(t, ErrBadCall, err)
	})

	t.Run("setup and release", func(t *testing.T) {
		_, present := os.LookupEnv("SAFE_TEST_A")
		require.False(t, present)

		release, err := SetEnvs("SAFE_TEST_A", "-")
		require.NoError(t, err)
		_, present = os.LookupEnv("SAFE_TEST_A")
		assert.True(t, present)

		require.NotPanics(t, func() { release(nil) })
		_, present = os.LookupEnv("SAFE_TEST_A")
		assert.False(t, present)
	})

	t.Run("setup present key", func(t *testing.T) {
		_, present := os.LookupEnv("SAFE_TEST_B")
		require.False(t, present)

		require.NoError(t, os.Setenv("SAFE_TEST_B", "-"))
		assert.Equal(t, "-", os.Getenv("SAFE_TEST_B"))

		release, err := SetEnvs("SAFE_TEST_B", "*")
		require.NoError(t, err)
		assert.Equal(t, "*", os.Getenv("SAFE_TEST_B"))

		require.NotPanics(t, func() { release(nil) })
		assert.Equal(t, "-", os.Getenv("SAFE_TEST_B"))
		require.NoError(t, os.Unsetenv("SAFE_TEST_B"))
	})

	t.Run("setup with errors", func(t *testing.T) {
		release, err := SetEnvs("", "value")
		assert.Error(t, err)
		assert.Nil(t, release)
	})
}
