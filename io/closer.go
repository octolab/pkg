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
	previous Closer
}

// Close implements io.Closer interface and calls a Close method on the chain.
func (closer cascadeCloser) Close() error {
	if closer.WriteCloser != nil {
		if err := closer.WriteCloser.Close(); err != nil {
			return err
		}
	}
	if err := closer.previous.Close(); err != nil {
		return err
	}
	if closer.ReadCloser != nil {
		return closer.ReadCloser.Close()
	}
	return nil
}
