package errors_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "go.octolab.org/errors"
)

const expected Message = "test"

func TestMessage_Error(t *testing.T) {
	caller := func() error {
		return Message("test")
	}
	assert.Equal(t, expected, caller())
	assert.EqualError(t, caller(), "test")
	assert.True(t, expected == caller())
}
