package errors_test

import (
	stderrors "errors"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	. "go.octolab.org/errors"
)

const expected Message = "error"

func TestMessage(t *testing.T) {
	caller := func() error {
		return Message("error")
	}
	assert.Equal(t, expected, caller())
	assert.EqualError(t, caller(), expected.Error())
	assert.True(t, expected == caller())
}

func TestUnwrap(t *testing.T) {
	cause := stderrors.New("origin")
	err := errors.WithMessage(cause, "wrapper")
	assert.NotEqual(t, cause, err)
	assert.Equal(t, cause, Unwrap(err))
}
