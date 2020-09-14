package encoding

import "io"

type Encoder interface {
	Encode(interface{}) error
}

type EncodeCloser interface {
	Encoder
	io.Closer
}

func ToEncodeCloser(encoder Encoder, wc io.WriteCloser) EncodeCloser {
	return struct {
		Encoder
		io.Closer
	}{encoder, wc}
}
