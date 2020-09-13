package io

import "io"

type ReadCloserChain func(io.ReadCloser) (io.ReadCloser, error)

func (wrapper ReadCloserChain) Wrap(wrapped ReadCloserChain) ReadCloserChain {
	return func(rc io.ReadCloser) (io.ReadCloser, error) {
		var err error
		rc, err = wrapped(rc)
		if err != nil {
			return nil, err
		}
		return wrapper(rc)
	}
}

// CascadeReadCloser returns a combination of two io.ReadCloser.
func CascadeReadCloser(current, previous io.ReadCloser) io.ReadCloser {
	return cascadeCloser{ReadCloser: current, previous: previous}
}

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
