package io

import "io"

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
