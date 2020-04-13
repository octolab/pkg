package safe_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/safe"
)

func TestClose(t *testing.T) {
	t.Run("with error", func(t *testing.T) {
		var called bool
		fn := (Closer)(func() error { return errors.New("test") })
		Close(fn, func(err error) { called = assert.Error(t, err) })
		assert.True(t, called)
	})
	t.Run("without error", func(t *testing.T) {
		var called bool
		fn := (Closer)(func() error { return nil })
		Close(fn, func(error) { called = true })
		assert.False(t, called)
	})
}
