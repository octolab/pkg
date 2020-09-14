package encoding

import "io"

type Decoder interface {
	Decode(interface{}) error
}

type DecodeCloser interface {
	Decoder
	io.Closer
}
