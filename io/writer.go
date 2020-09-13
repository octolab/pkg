package io

import "io"

// WriterNopCloser returns a io.WriteCloser with a no-op Close method
// wrapping the provided io.Writer.
func WriterNopCloser(writer io.Writer) io.WriteCloser {
	return nopCloser{Writer: writer}
}
