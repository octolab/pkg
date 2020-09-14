package encoding

import "io"

type Encoder interface {
	Encode(interface{}) error
}

type EncodeCloser interface {
	Encoder
	io.Closer
}
