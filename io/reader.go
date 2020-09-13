package io

import "io"

// ReadNopCloser returns a io.ReadCloser with a no-op Close method
// wrapping the provided io.Reader.
func ReadNopCloser(reader io.Reader) io.ReadCloser {
	return nopCloser{Reader: reader}
}

// ToReadCloser converts a io.Reader into io.ReadCloser.
func ToReadCloser(reader io.Reader) io.ReadCloser {
	if rc, is := reader.(io.ReadCloser); is {
		return rc
	}
	return ReadNopCloser(reader)
}
