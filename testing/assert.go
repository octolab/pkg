package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// NoError returns a helper as an error handler
// to check no error occurred.
func NoError(t *testing.T) func(error) {
	return func(err error) {
		t.Helper()
		assert.NoError(t, err)
	}
}

// StrictNoError returns a helper as a strict
// error handler to check no error occurred.
func StrictNoError(t *testing.T) func(error) {
	return func(err error) {
		t.Helper()
		require.NoError(t, err)
	}
}
