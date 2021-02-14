package env_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "go.octolab.org/env"
	helper "go.octolab.org/testing"
)

func TestEnvironment(t *testing.T) {
	release, err := helper.SetEnvs(
		helper.StrictNoError(t),
		Must("KEY", "v1.0"),
		Must("Key", "v2.0"),
	)
	require.NoError(t, err)

	env := From(os.Environ())
	assert.Equal(t, os.Environ(), env.Environ())

	t.Run("lookup", func(t *testing.T) {
		val, has := env.Lookup("KEY")
		assert.True(t, has)
		assert.Equal(t, "KEY", val.Name())
		assert.Equal(t, "v1.0", val.Value())
	})

	t.Run("case-sensitive", func(t *testing.T) {
		val, has := env.Lookup("Key")
		assert.True(t, has)
		assert.Equal(t, "Key", val.Name())
		assert.Equal(t, "v2.0", val.Value())
	})

	t.Run("no exists", func(t *testing.T) {
		val, has := env.Lookup("KeY")
		assert.False(t, has)
		assert.Empty(t, val.Name())
		assert.Empty(t, val.Value())
	})

	t.Run("bad input", func(t *testing.T) {
		vars := From([]string{
			"KEY1=v1.0",
			"KEY2=",
			"KEY3",
		})

		for _, test := range []struct {
			key string
			val string
			has bool
		}{
			{"KEY1", "v1.0", true},
			{"KEY2", "", true},
			{"KEY3", "", false},
		} {
			val, has := vars.Lookup(test.key)
			assert.Equal(t, test.has, has)
			assert.Equal(t, test.val, val.Value())
		}
	})

	release(helper.StrictNoError(t))
}
