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
	os.Clearenv()
	vars := []struct {
		env.Variable
		present bool
		replace string
	}{
		{env.Must("KEY1", "v1.0"), true, "v2.0"},
		{env.Must("KEY2", ""), true, "v1.0"},
		{env.Must("KEY3", ""), false, "single"},
	}
	for _, v := range vars {
		if v.present {
			require.NoError(t, os.Setenv(v.Key(), v.Value()))
		}
	}

	var (
		faker   = gofakeit.New(time.Now().UnixNano())
		noError = func(err error) { assert.NoError(t, err) }
		verify  = func() {
			for _, v := range vars {
				val, present := os.LookupEnv(v.Key())
				assert.Equal(t, v.present, present)
				assert.Equal(t, v.Value(), val)
			}
		}
	)

	t.Run("no variables", func(t *testing.T) {
		release, err := SetEnvs(noError)
		require.NoError(t, err)
		assert.NotPanics(t, func() { release(nil) })
	})

	t.Run("setup and release", func(t *testing.T) {
		verify()

		vv := make([]env.Variable, 0, len(vars))
		for _, v := range vars {
			vv = append(vv, env.Must(v.Key(), v.replace))
		}
		release, err := SetEnvs(noError, vv...)
		require.NoError(t, err)
		for _, v := range vars {
			val, present := os.LookupEnv(v.Key())
			assert.True(t, present)
			assert.Equal(t, v.replace, val)
		}

		release(noError)
		verify()
	})

	t.Run("parallel run", func(t *testing.T) {
		for i := range sequence.Simple(10) {
			t.Run(strconv.Itoa(i), func(t *testing.T) {
				t.Parallel()

				vv := make([]env.Variable, 0, len(vars))
				for _, v := range vars {
					vv = append(vv, env.Must(v.Key(), faker.AppVersion()))
				}
				release, err := SetEnvs(noError, vv...)
				require.NoError(t, err)
				for i, v := range vars {
					val, present := os.LookupEnv(v.Key())
					assert.True(t, present)
					assert.Equal(t, vv[i].Value(), val)
				}

				release(noError)
			})
		}
	})

	verify()
	os.Clearenv()
}
