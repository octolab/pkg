// +build go1.13

package errors_test

import (
	stderrors "errors"
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	. "go.octolab.org/errors"
)

func TestUnwrap_Builtin(t *testing.T) {
	t.Run("compatible with built-in errors", func(t *testing.T) {
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
