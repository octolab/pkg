// +build go1.13

package errors_test

import (
	"net"
	"os"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	. "go.octolab.org/errors"
)

func TestClassifier_Classify(t *testing.T) {
	classifier := make(Classifier).
		ClassifyAs(networkClass, new(NetworkError)).
		ClassifyAs(filesystemClass, os.ErrExist, os.ErrNotExist)

	tests := map[string]struct {
		err      error
		expected string
	}{
		"naked error":   {os.ErrExist, filesystemClass},
		"wrapped error": {errors.Wrap(os.ErrExist, "wrapped"), filesystemClass},
		"network error": {new(net.AddrError), networkClass},
		"custom error":  {new(network), networkClass},
		"nil error":     {nil, Unknown},
		"unknown error": {os.ErrInvalid, Unknown},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, classifier.Classify(test.err, Unknown))
		})
	}
}

func TestClassifier_Consistent(t *testing.T) {
	tests := map[string]struct {
		classifier Classifier
		expected   bool
	}{
		"consistent": {
			make(Classifier).
				ClassifyAs(networkClass, new(NetworkError)).
				ClassifyAs(filesystemClass, os.ErrExist, os.ErrNotExist),
			true,
		},
		"not consistent": {
			make(Classifier).
				ClassifyAs(networkClass, new(NetworkError)).
				ClassifyAs(timeoutClass, new(TimeoutError), new(NetworkError)),
			true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.classifier.Consistent())
		})
	}
}

// helper

const (
	networkClass    = "network"
	filesystemClass = "fs"
	timeoutClass    = "timeout"
)

type network struct{}

func (err *network) Error() string   { return "custom" }
func (err *network) Timeout() bool   { return true }
func (err *network) Temporary() bool { return true }
