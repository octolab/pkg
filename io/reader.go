package io

// CascadeReadCloser returns a combination of two io.ReadCloser.
func CascadeReadCloser(current, previous ReadCloser) ReadCloser {
	return cascadeCloser{ReadCloser: current, previous: previous}
}

// ReadNopCloser returns a io.ReadCloser with a no-op Close method
// wrapping the provided io.Reader.
func ReadNopCloser(reader Reader) ReadCloser {
	return nopCloser{Reader: reader}
}

// ToReadCloser converts a io.Reader into io.ReadCloser.
func ToReadCloser(reader Reader) ReadCloser {
	if rc, is := reader.(ReadCloser); is {
		return rc
	}
	return ReadNopCloser(reader)
}

type ReadCloserChain func(ReadCloser) (ReadCloser, error)

func (head ReadCloserChain) Chain(link ReadCloserChain) ReadCloserChain {
	return func(rc ReadCloser) (ReadCloser, error) {
		var err error
		rc, err = link(rc)
		if err != nil {
			return nil, err
		}
		return head(rc)
	}
}
