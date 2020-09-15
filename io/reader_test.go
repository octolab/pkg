package io_test

import (
	"bytes"
	"compress/gzip"
	"compress/lzw"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "go.octolab.org/io"
	"go.octolab.org/unsafe"
)

func TestCascadeReadCloser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("lifo check", func(t *testing.T) {
		var chain []string

		rc1 := NewMockReadCloser(ctrl)
		rc1.EXPECT().
			Close().
			Return(nil).
			Do(func() { chain = append(chain, "rc1") })
		rc2 := NewMockReadCloser(ctrl)
		rc2.EXPECT().
			Close().
			Return(nil).
			Do(func() { chain = append(chain, "rc2") })

		lifo := CascadeReadCloser(rc1, rc2)
		require.NoError(t, lifo.Close())
		assert.Equal(t, []string{"rc2", "rc1"}, chain)
	})

	t.Run("error on chain", func(t *testing.T) {
		rc1 := NewMockReadCloser(ctrl)
		rc1.EXPECT().
			Close().
			Return(nil).
			Times(0)
		rc2 := NewMockReadCloser(ctrl)
		rc2.EXPECT().
			Close().
			Return(errors.New("cannot close"))

		lifo := CascadeReadCloser(rc1, rc2)
		require.Error(t, lifo.Close())
	})
}

func ExampleCascadeReadCloser() {
	input := bytes.NewReader([]byte{
		31, 139, 8, 0, 0, 0, 0, 0, 0, 255, 202, 125, 54, 135,
		147, 3, 16, 0, 0, 255, 255, 234, 40, 132, 127, 5, 0, 0, 0,
	})

	last, err := gzip.NewReader(input)
	if err != nil {
		panic(err)
	}
	first := lzw.NewReader(last, lzw.LSB, 8)
	cascade := CascadeReadCloser(first, last)

	buf := bytes.NewBuffer(nil)
	unsafe.DoSilent(io.Copy(buf, cascade))
	if err := cascade.Close(); err != nil {
		panic(err)
	}
	// it calls first last.Close() and then first.Close()

	fmt.Println(buf.String())
	// output: msg
}
