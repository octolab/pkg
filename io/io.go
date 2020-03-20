package io

import (
	"bytes"
	"io"
	"io/ioutil"
)

func RepeatableReadCloser(body io.ReadCloser, buf *bytes.Buffer) io.ReadCloser {
	return &repeatable{src: TeeReadCloser(body, buf), dst: buf}
}

type repeatable struct {
	src io.ReadCloser
	dst *bytes.Buffer
}

func (r *repeatable) Read(p []byte) (n int, err error) {
	n, err = r.src.Read(p)
	if n < len(p) && (err == nil || err == io.EOF) {
		buf := bytes.NewBuffer(make([]byte, 0, r.dst.Len()))
		r.src, r.dst = TeeReadCloser(ioutil.NopCloser(r.dst), buf), buf
		err = io.EOF
	}
	return
}

func (r *repeatable) Close() error { return r.src.Close() }

// TeeReadCloser returns a ReadCloser that writes to w what it reads from rc.
// All reads from rc performed through it are matched with
// corresponding writes to w. There is no internal buffering -
// the write must complete before the read completes.
// Any error encountered while writing is reported as a read error.
//
//  func Handler(rw http.ResponseWriter, req *http.Request) {
//  	buf := bytes.NewBuffer(make([]byte, 0, req.ContentLength))
//  	body := io.TeeReadCloser(req.Body, buf)
//
//  	var payload interface{}
//  	if err := json.NewDecoder(body).Decode(&payload); err != nil {
//  		log.Printf("invalid json: %q", buf.String())
//  	}
//  }
//
func TeeReadCloser(rc io.ReadCloser, w io.Writer) io.ReadCloser {
	type pipe struct {
		io.Reader
		io.Closer
	}
	return pipe{io.TeeReader(rc, w), rc}
}
