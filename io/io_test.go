package io_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "go.octolab.org/io"
	"go.octolab.org/safe"
	"go.octolab.org/unsafe"
)

func TestRepeatableReadCloser(t *testing.T) {
	t.Run("repeatable request", func(t *testing.T) {
		payload := `{"payload": "test"}`
		req, err := http.NewRequest(
			http.MethodPost, "/",
			RepeatableReadCloser(
				ioutil.NopCloser(strings.NewReader(payload)),
				bytes.NewBuffer(make([]byte, 0, len(payload))),
			),
		)
		require.NoError(t, err)

		var handler http.HandlerFunc = func(rw http.ResponseWriter, req *http.Request) {
			var obtained struct {
				Payload string `json:"payload"`
			}
			require.NoError(t, json.NewDecoder(req.Body).Decode(&obtained))
			require.NoError(t, req.Body.Close())
			require.Equal(t, "test", obtained.Payload)
		}

		for i := range make([]struct{}, 10) {
			iteration := fmt.Sprintf("iteration #%d", i+1)
			t.Run(iteration, func(t *testing.T) {
				handler.ServeHTTP(httptest.NewRecorder(), req)
			})
		}
	})
	t.Run("repeatable read", func(t *testing.T) {
		tests := map[string]struct {
			source   io.ReadCloser
			expected string
		}{
			"strings": {
				source: RepeatableReadCloser(
					ioutil.NopCloser(strings.NewReader("test")),
					bytes.NewBuffer(make([]byte, 0, 4)),
				),
				expected: "test",
			},
			"bytes": {
				source: RepeatableReadCloser(
					ioutil.NopCloser(bytes.NewReader([]byte("test"))),
					bytes.NewBuffer(make([]byte, 0, 4)),
				),
				expected: "test",
			},
			"file": {
				source: RepeatableReadCloser(
					file("./testdata/test.txt"),
					bytes.NewBuffer(make([]byte, 0, 5)),
				),
				expected: "test\n",
			},
		}

		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				buf := bytes.NewBuffer(make([]byte, 0, len(test.expected)/2))
				for i := range make([]struct{}, 10) {
					iteration := fmt.Sprintf("iteration #%d", i+1)
					n, err := buf.ReadFrom(test.source)
					require.NoError(t, err, iteration)
					require.Equal(t, int64(buf.Len()), n, iteration)
					require.Equal(t, test.expected, buf.String(), iteration)
					require.NoError(t, test.source.Close(), iteration)
					buf.Reset()
				}
			})
		}
	})
}

func ExampleRepeatableReadCloser() {
	i, responses := int32(-1), []int{
		http.StatusInternalServerError,
		http.StatusServiceUnavailable,
		http.StatusOK,
	}

	var (
		handler http.HandlerFunc = func(rw http.ResponseWriter, req *http.Request) {
			unsafe.DoSilent(io.Copy(ioutil.Discard, req.Body))
			unsafe.Ignore(req.Body.Close())

			code := responses[atomic.AddInt32(&i, 1)]
			http.Error(rw, http.StatusText(code), code)
		}

		request = http.Request{
			Method: http.MethodPost,
			Body: RepeatableReadCloser(
				ioutil.NopCloser(strings.NewReader("echo")),
				bytes.NewBuffer(make([]byte, 0, 4)),
			),
		}
	)

	for {
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, &request)
		unsafe.DoSilent(io.Copy(os.Stdout, recorder.Body))
		if recorder.Code == http.StatusOK {
			break
		}
	}
	// output:
	// Internal Server Error
	// Service Unavailable
	// OK
}

func TestTeeReadCloser(t *testing.T) {
	payload := "invalid json"
	req, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	require.NoError(t, err)
	req = req.WithContext(context.Background())

	var handler http.HandlerFunc = func(rw http.ResponseWriter, req *http.Request) {
		buf := bytes.NewBuffer(make([]byte, 0, req.ContentLength))
		body := TeeReadCloser(req.Body, buf)

		var expected []int
		assert.Error(t, json.NewDecoder(body).Decode(&expected))
		assert.NoError(t, body.Close())
		assert.Nil(t, expected)
		assert.Equal(t, payload, buf.String())
	}
	handler.ServeHTTP(httptest.NewRecorder(), req)
}

func ExampleTeeReadCloser() {
	var (
		handler http.HandlerFunc = func(rw http.ResponseWriter, req *http.Request) {
			buf := bytes.NewBuffer(make([]byte, 0, req.ContentLength))
			body := TeeReadCloser(req.Body, buf)
			defer safe.Close(body, unsafe.Ignore)

			var payload interface{}
			if err := json.NewDecoder(body).Decode(&payload); err != nil {
				message := fmt.Sprintf("invalid json: %s", buf.String())
				http.Error(rw, message, http.StatusBadRequest)
			}
		}

		request = http.Request{
			Method: http.MethodPost,
			Body:   ioutil.NopCloser(strings.NewReader(`{bad: "json"}`)),
		}
	)

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, &request)
	unsafe.DoSilent(io.Copy(os.Stdout, recorder.Body))
	// output:
	// invalid json: {bad: "json"}
}

// helpers

func file(name string) *os.File {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	return f
}
