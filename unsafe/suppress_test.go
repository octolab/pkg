package unsafe_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/unsafe"
)

func TestDoSilent(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	DoSilent(fmt.Fprintf(buf, "test"))
	assert.Equal(t, "test", buf.String())

	to, from := bytes.NewBuffer(nil), strings.NewReader("test")
	DoSilent(io.Copy(to, from))
	assert.Equal(t, "test", to.String())
}

func TestIgnore(t *testing.T) {
	var data map[string]interface{}
	Ignore(json.NewDecoder(strings.NewReader(`{5: 12}`)).Decode(&data))
	assert.Nil(t, data)
}
