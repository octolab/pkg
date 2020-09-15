package io

type nopCloser struct {
	Reader
	Writer
}

// Close implements io.Closer interface and do nothing.
func (nopCloser) Close() error { return nil }

type cascadeCloser struct {
	ReadCloser
	WriteCloser
	close func() error
}

// Close implements io.Closer interface and calls a Close method on the chain.
func (closer cascadeCloser) Close() error { return closer.close() }
