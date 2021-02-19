package io

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
)

//go:generate mockgen -destination mocks_test.go -package ${GOPACKAGE}_test io Closer,Reader,ReadCloser,Writer,WriteCloser

// Aliases to make the package more self-sufficient.
type (
	Closer      = io.Closer
	Reader      = io.Reader
	ReadCloser  = io.ReadCloser
	Writer      = io.Writer
	WriteCloser = io.WriteCloser
)

func Discard(body ReadCloser) Closer {
	if _, err := io.Copy(ioutil.Discard, body); err != nil {
		return closer(func() error { return err })
	}
	return body
}

// RepeatableReadCloser returns a ReadCloser that can be read an unlimited number of times.
//
//  payload := strings.NewReader(`{"some":"payload"}`)
//  body := RepeatableReadCloser(
//  	ioutil.NopCloser(payload),
//  	bytes.NewBuffer(make([]byte, 0, payload.Len())),
//  )
//  req, err := http.NewRequest(http.MethodPost, "/api", body)
//  if err != nil {
//  	log.Fatal(err)
//  }
//  for {
//  	resp, err := http.DefaultClient.Do(req)
//  	if err == nil && resp.StatusCode == http.StatusOK {
//  		break
//  	}
//  	time.Sleep(time.Second)
//  }
//
func RepeatableReadCloser(body ReadCloser, buf *bytes.Buffer) ReadCloser {
	return &repeatable{src: TeeReadCloser(body, buf), dst: buf}
}

// TeeReadCloser returns a ReadCloser that writes to w what it reads from rc.
// All reads from rc performed through it are matched with
// corresponding writes to w. There is no internal buffering -
// the write must complete before the read completes.
// Any error encountered while writing is reported as a read error.
//
//  func Handler(rw http.ResponseWriter, req *http.Request) {
//  	buf := bytes.NewBuffer(make([]byte, 0, req.ContentLength))
//  	body := io.TeeReadCloser(req.Body, buf)
//  	defer safe.Close(body, unsafe.Ignore)
//
//  	var payload interface{}
//  	if err := json.NewDecoder(body).Decode(&payload); err != nil {
//  		message := fmt.Sprintf("invalid json: %s", buf.String())
//  		http.Error(rw, message, http.StatusBadRequest)
//  	}
//  }
//
func TeeReadCloser(rc ReadCloser, w Writer) ReadCloser {
	type pipe struct {
		Reader
		Closer
	}
	return pipe{io.TeeReader(rc, w), rc}
}

type repeatable struct {
	src ReadCloser
	dst *bytes.Buffer
}

// Read implements the io.Reader interface.
// It reads from the underlying TeeReadCloser
// and rotates it if it is done.
func (r *repeatable) Read(p []byte) (n int, err error) {
	n, err = r.src.Read(p)
	var eof bool
	eof = n == 0 && errors.Is(err, io.EOF)
	eof = eof || (n < len(p) && err == nil) // danger zone ("repeatable request")
	if eof {
		buf := bytes.NewBuffer(make([]byte, 0, r.dst.Len()))
		r.src, r.dst = TeeReadCloser(ioutil.NopCloser(r.dst), buf), buf
		err = io.EOF // prevent infinite loop related to danger zone ("repeatable read")
	}
	return
}

// Close implements the io.Closer interface.
// It closes the underlying ReadCloser.
func (r *repeatable) Close() error { return r.src.Close() }
