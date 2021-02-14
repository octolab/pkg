package testing_test

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.octolab.org/env"
	"go.octolab.org/sequence"
	. "go.octolab.org/testing"
)

func TestSetEnvs(t *testing.T) {
	set := []struct {
		env.Variable
		present bool
		replace string
	}{
		{env.Must("KEY1", "v1.0"), true, "v2.0"},
		{env.Must("KEY2", ""), true, "v1.0"},
		{env.Must("KEY3", ""), false, "single"},
	}
	for _, v := range set {
		if v.present {
			require.NoError(t, os.Setenv(v.Name(), v.Value()))
		}
	}
	verify := func() {
		t.Helper()
		for _, v := range set {
			val, present := os.LookupEnv(v.Name())
			assert.Equal(t, v.present, present)
			assert.Equal(t, v.Value(), val)
		}
	}

	t.Run("no variables", func(t *testing.T) {
		release, err := SetEnvs(NoError(t))
		require.NoError(t, err)
		assert.NotPanics(t, func() { release(nil) })
	})

	t.Run("setup and release", func(t *testing.T) {
		verify()

		vars := make([]env.Variable, 0, len(set))
		for _, v := range set {
			vars = append(vars, env.Must(v.Name(), v.replace))
		}
		release, err := SetEnvs(NoError(t), vars...)
		require.NoError(t, err)
		for _, v := range set {
			val, present := os.LookupEnv(v.Name())
			assert.True(t, present)
			assert.Equal(t, v.replace, val)
		}

		release(StrictNoError(t))
		verify()
	})

	t.Run("parallel run", func(t *testing.T) {
		for i := range sequence.Simple(10) {
			t.Run(strconv.Itoa(i), func(t *testing.T) {
				t.Parallel()

				fake := gofakeit.New(time.Now().UnixNano())
				vars := make([]env.Variable, 0, len(set))
				for _, v := range set {
					vars = append(vars, env.Must(v.Name(), fake.AppVersion()))
				}
				release, err := SetEnvs(NoError(t), vars...)
				require.NoError(t, err)
				for i, v := range set {
					val, present := os.LookupEnv(v.Name())
					assert.True(t, present)
					assert.Equal(t, vars[i].Value(), val)
				}

				release(NoError(t))
			})
		}
	})

	verify()
	for _, v := range set {
		if v.present {
			require.NoError(t, os.Unsetenv(v.Name()))
		}
	}
}
