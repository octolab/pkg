package io_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	. "go.octolab.org/io"
)

func TestTeeReadCloser(t *testing.T) {
	payload := "invalid json"
	handler := func(rw http.ResponseWriter, req *http.Request) {
		buf := bytes.NewBuffer(nil)
		body := TeeReadCloser(req.Body, buf)

		var expected []int
		assert.Error(t, json.NewDecoder(body).Decode(&expected))
		assert.Nil(t, expected)
		assert.Equal(t, payload, buf.String())
	}

	rw := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	assert.NoError(t, err)

	handler(rw, req)
}
