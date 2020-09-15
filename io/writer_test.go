package io_test

import (
	"bytes"
	"compress/gzip"
	"compress/lzw"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "go.octolab.org/io"
	"go.octolab.org/unsafe"
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

func ExampleCascadeWriteCloser() {
	buf, msg := bytes.NewBuffer(nil), "msg"

	last := gzip.NewWriter(buf)
	first := lzw.NewWriter(last, lzw.LSB, 8)
	cascade := CascadeWriteCloser(first, last)

	unsafe.DoSilent(io.Copy(cascade, strings.NewReader(msg)))
	if err := cascade.Close(); err != nil {
		panic(err)
	}
	// it calls first first.Close() and then last.Close()

	fmt.Println(buf.Bytes())
	// output: [31 139 8 0 0 0 0 0 0 255 202 125 54 135 147 3 16 0 0 255 255 234 40 132 127 5 0 0 0]
}
