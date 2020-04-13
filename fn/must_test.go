package fn_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/fn"
)

func TestMust(t *testing.T) {
	tests := map[string]struct {
		actions []func() error
		assert  func(assert.TestingT, assert.PanicTestFunc, ...interface{}) bool
	}{
		"with panic": {
			[]func() error{
				func() error { return nil },
				func() error { return errors.New("raise panic") },
				func() error { return nil },
			},
			assert.Panics,
		},
		"without panic": {
			[]func() error{
				func() error { return nil },
				func() error { return nil },
				func() error { return nil },
			},
			assert.NotPanics,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.assert(t, func() { Must(test.actions...) })
		})
	}
}
