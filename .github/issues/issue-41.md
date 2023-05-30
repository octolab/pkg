---
code: null
id: MDU6SXNzdWU4MDk4OTExNTY=
databaseId: 809891156
number: 41
url: https://github.com/octolab/pkg/issues/41
title: "request builder for testing"
issueType: null
labels:
  - "help wanted"
assignees:
  - kamilsk
milestone: null
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2021-02-17T06:15:53Z
updatedAt: 2021-02-17T06:33:42Z
lastEditedAt: null
closedAt: 2021-02-17T06:33:42Z
---

# request builder for testing

```go
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
```
