package io

// CascadeReadCloser combines two io.ReadCloser's Close methods
// into one LIFO chain and returns first of them as a result io.ReadCloser.
//
//	func DecompressInput(input []byte) ([]byte, error) {
//		last, err := gzip.NewReader(bytes.NewReader(input))
//		if err != nil {
//			return nil, err
//		}
//		first := lzw.NewReader(last, lzw.LSB, 8)
//		cascade := CascadeReadCloser(first, last)
//
//		buf := bytes.NewBuffer(make([]byte, 0, len(input)))
//		if _, err := io.Copy(buf, cascade); err != nil {
//			return nil, err
//		}
//		return buf.Bytes(), cascade.Close() // it calls first last.Close() and then first.Close()
//	}
func CascadeReadCloser(current, previous ReadCloser) ReadCloser {
	return cascadeCloser{
		ReadCloser: current,
		close: func() error {
			if err := previous.Close(); err != nil {
				return err
			}
			return current.Close()
		},
	}
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
