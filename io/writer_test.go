package io_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "go.octolab.org/io"
)

func TestCascadeWriteCloser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("fifo check", func(t *testing.T) {
		var chain []string

		wc1 := NewMockWriteCloser(ctrl)
		wc1.EXPECT().
			Close().
			Return(nil).
			Do(func() { chain = append(chain, "wc1") })
		wc2 := NewMockWriteCloser(ctrl)
		wc2.EXPECT().
			Close().
			Return(nil).
			Do(func() { chain = append(chain, "wc2") })

		fifo := CascadeWriteCloser(wc1, wc2)
		require.NoError(t, fifo.Close())
		assert.Equal(t, []string{"wc1", "wc2"}, chain)
	})

	t.Run("error on chain", func(t *testing.T) {
		wc1 := NewMockWriteCloser(ctrl)
		wc1.EXPECT().
			Close().
			Return(errors.New("cannot close"))
		wc2 := NewMockWriteCloser(ctrl)
		wc2.EXPECT().
			Close().
			Return(nil).
			Times(0)

		fifo := CascadeWriteCloser(wc1, wc2)
		require.Error(t, fifo.Close())
	})
}
