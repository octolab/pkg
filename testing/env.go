package testing

import (
	"fmt"
	"os"
	"sync"

	"go.octolab.org/env"
)

// SetEnvs allows changing environment variables concurrently.
//
//  func Test(t *testing.T) {
//  	os.Clearenv()
//
//  	t.Run("case 1", func(t *testing.T) {
//  		t.Parallel()
//
//  		release, err := SetEnvs(NoError, env.Must(env.GoTraceback, "system"))
//  		require.NoError(t, err)
//
//  		assert.Equal(t, "system", os.Getenv(env.GoTraceback))
//  		release(StrictNoError)
//  	})
//
//  	t.Run("case 2", func(t *testing.T) {
//  		t.Parallel()
//
//  		release, err := SetEnvs(NoError, env.Must(env.GoTraceback, "crash"))
//  		require.NoError(t, err)
//
//  		assert.Equal(t, "crash", os.Getenv(env.GoTraceback))
//  		release(StrictNoError)
//  	})
//
//  	_, present := os.LookupEnv(env.GoTraceback)
//  	assert.False(t, present)
//  }
//
func SetEnvs(handle func(error), vars ...env.Variable) (func(func(error)), error) {
	return setEnvs(os.LookupEnv, os.Setenv, os.Unsetenv, handle, vars...)
}

var guard sync.Mutex

func setEnvs(
	lookup func(string) (string, bool),
	set func(string, string) error,
	unset func(string) error,
	handle func(error),
	vars ...env.Variable,
) (func(func(error)), error) {
	if len(vars) == 0 {
		return func(func(error)) {}, nil
	}

	var (
		err error
		pos = -1
	)
	before := make(map[string]string, len(vars))

	guard.Lock()
	for i, v := range vars {
		if val, present := lookup(v.Key()); present {
			before[v.Key()] = val
		}
		if err = set(v.Key(), v.Value()); err != nil {
			err = fmt.Errorf("cannot set environment variable %s: %w", v, err)
			break
		}
		pos = i
	}

	rollback := func(handle func(error)) {
		defer guard.Unlock()
		for _, v := range vars[:pos+1] {
			if prev, found := before[v.Key()]; found {
				handle(set(v.Key(), prev))
			} else {
				handle(unset(v.Key()))
			}
		}
	}

	if err != nil {
		rollback(handle)
		return nil, err
	}
	return rollback, nil
}
