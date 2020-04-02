package errors_test

import (
	stderrors "errors"
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	. "go.octolab.org/errors"
	"go.octolab.org/runtime"
)

const expected Message = "test"

func TestMessage(t *testing.T) {
	caller := func() error {
		return Message("test")
	}
	assert.Equal(t, expected, caller())
	assert.EqualError(t, caller(), "test")
	assert.True(t, expected == caller())
}

func TestUnwrap(t *testing.T) {
	t.Run("compatible with github.com/pkg/errors", func(t *testing.T) {
		cause := stderrors.New("origin")
		err := errors.WithMessage(cause, "wrapper")
		assert.NotEqual(t, cause, err)
		assert.Equal(t, cause, Unwrap(err))
	})
	t.Run("compatible with built-in errors", func(t *testing.T) {
		version := runtime.Version()
		if version.Minor < 13 {
			t.SkipNow()
		}
		cause := stderrors.New("origin")
		err := fmt.Errorf("wrapper: %w", cause)
		assert.NotEqual(t, cause, err)
		assert.Equal(t, cause, Unwrap(err))
	})
	t.Run("onion", func(t *testing.T) {
		cause := stderrors.New("origin")
		err := fmt.Errorf("wrapper: %w", errors.WithMessage(layer{cause}, "wrapper"))
		assert.NotEqual(t, cause, err)
		assert.Equal(t, cause, Unwrap(err))
	})
}

// helpers

type layer struct{ error }

func (onion layer) Cause() error { return onion.error }
