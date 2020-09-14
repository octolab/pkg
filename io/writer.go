package io

import "io"

// CascadeWriteCloser returns a combination of two io.WriteCloser.
func CascadeWriteCloser(current, previous io.WriteCloser) io.WriteCloser {
	return cascadeCloser{WriteCloser: current, previous: previous}
}

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

type WriteCloserChain func(io.WriteCloser) (io.WriteCloser, error)

func (wrapper WriteCloserChain) Wrap(wrapped WriteCloserChain) WriteCloserChain {
	return func(wc io.WriteCloser) (io.WriteCloser, error) {
		var err error
		wc, err = wrapped(wc)
		if err != nil {
			return nil, err
		}
		return wrapper(wc)
	}
}
