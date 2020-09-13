package io

import "io"

type nopCloser struct {
	io.Reader
	io.Writer
}

// Close implements io.Closer interface and do nothing.
func (nopCloser) Close() error { return nil }

type cascadeCloser struct {
	io.ReadCloser
	io.WriteCloser
	previous io.Closer
}

// Close implements io.Closer interface and calls a Close method on the chain.
func (closer cascadeCloser) Close() error {
	if err := closer.previous.Close(); err != nil {
		return err
	}
	if closer.ReadCloser != nil {
		return closer.ReadCloser.Close()
	}
	return closer.WriteCloser.Close()
}
