package safe_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "go.octolab.org/errors"
	. "go.octolab.org/safe"
)

func TestDo(t *testing.T) {
	tests := map[string]struct {
		action  func() error
		handler func(error)
	}{
		"with error": {
			func() error { return errors.New("error") },
			func(err error) { assert.EqualError(t, err, "error") },
		},
		"with panic": {
			func() error { panic("test") },
			func(err error) {
				recovered, is := Unwrap(err).(Recovered)
				require.True(t, is)
				require.NotNil(t, recovered)
				assert.Equal(t, "unexpected panic occurred", recovered.Error())
				assert.Equal(t, "test", recovered.Cause())
			},
		},
		"without anything": {
			func() error { return nil },
			func(error) { t.Fail() },
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.NotPanics(t, func() { Do(test.action, test.handler) })
		})
	}
}
