package io

// CascadeWriteCloser returns a combination of two io.WriteCloser.
func CascadeWriteCloser(current, previous WriteCloser) WriteCloser {
	return cascadeCloser{WriteCloser: current, previous: previous}
}

// WriteNopCloser returns a io.WriteCloser with a no-op Close method
// wrapping the provided io.Writer.
func WriteNopCloser(writer Writer) WriteCloser {
	return nopCloser{Writer: writer}
}

// ToWriteCloser converts a io.Writer into io.WriteCloser.
func ToWriteCloser(writer Writer) WriteCloser {
	if wc, is := writer.(WriteCloser); is {
		return wc
	}
	return WriteNopCloser(writer)
}

type WriteCloserChain func(WriteCloser) (WriteCloser, error)

func (head WriteCloserChain) Chain(link WriteCloserChain) WriteCloserChain {
	return func(wc WriteCloser) (WriteCloser, error) {
		var err error
		wc, err = link(wc)
		if err != nil {
			return nil, err
		}
		return head(wc)
	}
}
