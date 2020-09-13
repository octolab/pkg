package io

import "io"

// WriteNopCloser returns a io.WriteCloser with a no-op Close method
// wrapping the provided io.Writer.
func WriteNopCloser(writer io.Writer) io.WriteCloser {
	return nopCloser{Writer: writer}
}

// ToWriteCloser converts a io.Writer into io.WriteCloser.
func ToWriteCloser(writer io.Writer) io.WriteCloser {
	if wc, is := writer.(io.WriteCloser); is {
		return wc
	}
	return WriteNopCloser(writer)
}
