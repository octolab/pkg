package io

import "io"

type nopCloser struct {
	io.Reader
	io.Writer
}

// Close implements io.Closer interface.
func (nopCloser) Close() error { return nil }
