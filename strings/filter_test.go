package strings_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/strings"
)

func TestFirstValid(t *testing.T) {
	tests := map[string]struct {
		strings  []string
		expected string
	}{
		"nothing to pass": {},
		"simple usage":    {[]string{"", "", "third", "fourth"}, "third"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, FirstNotEmpty(test.strings...))
		})
	}
}

func TestNotEmpty(t *testing.T) {
	tests := map[string]struct {
		strings  []string
		expected []string
	}{
		"nothing to pass": {},
		"simple usage": {
			[]string{"one", "", "two", "", "three"},
			[]string{"one", "two", "three"},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, NotEmpty(test.strings))
		})
	}
}
