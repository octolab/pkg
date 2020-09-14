package encoding

import "io"

type Decoder interface {
	Decode(interface{}) error
}

type DecodeCloser interface {
	Decoder
	io.Closer
}

func ToDecodeCloser(decoder Decoder, rc io.ReadCloser) DecodeCloser {
	return struct {
		Decoder
		io.Closer
	}{decoder, rc}
}
