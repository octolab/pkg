package http

import (
	"net/http"
	"net/url"
)

func NewRequest(options ...func(*http.Request)) *http.Request {
	req := new(http.Request)
	for _, configure := range options {
		configure(req)
	}
	return req
}

func WithForm(kv ...string) func(*http.Request) {
	return func(req *http.Request) {
		req.Form = make(url.Values, len(kv))

		for i := 0; i+1 < len(kv); i += 2 {
			req.Form.Set(kv[i], kv[i+1])
		}
	}
}

func WithHeaders(kv ...string) func(*http.Request) {
	return func(req *http.Request) {
		req.Header = make(http.Header, len(kv))

		for i := 0; i+1 < len(kv); i += 2 {
			req.Header.Set(kv[i], kv[i+1])
		}
	}
}

func WithPath(path string) func(*http.Request) {
	return func(req *http.Request) {
		req.URL = new(url.URL)
		req.URL.Path = path
	}
}
