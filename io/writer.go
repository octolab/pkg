package io

// CascadeWriteCloser combines two io.WriteCloser's Close methods
// into one FIFO chain and returns first of them as a result io.WriteCloser.
//
//	func CompressInput(input []byte) ([]byte, error) {
//		buf := bytes.NewBuffer(make([]byte, 0, len(input)))
//
//		last := gzip.NewWriter(buf)
//		first := lzw.NewWriter(last, lzw.LSB, 8)
//		cascade := CascadeWriteCloser(first, last)
//
//		if _, err := io.Copy(cascade, bytes.NewReader(input)); err != nil {
//			return nil, err
//		}
//
//		return buf.Bytes(), cascade.Close() // it calls first first.Close() and then last.Close()
//	}
func CascadeWriteCloser(current, previous WriteCloser) WriteCloser {
	return cascadeCloser{
		WriteCloser: current,
		close: func() error {
			if err := current.Close(); err != nil {
				return err
			}
			return previous.Close()
		},
	}
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
