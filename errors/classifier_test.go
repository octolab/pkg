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
		ClassifyAs(fatalClass, new(RecoveredError)).
		ClassifyAs(filesystemClass, os.ErrExist, os.ErrNotExist).
		ClassifyAs(repeatableClass, new(RetriableError)).
		ClassifyAs(Unknown, nil)

	tests := map[string]struct {
		err      error
		expected string
	}{
		"naked error":   {os.ErrExist, filesystemClass},
		"wrapped error": {errors.Wrap(os.ErrExist, "wrapped"), filesystemClass},
		"network error": {new(net.AddrError), networkClass},
		"custom error":  {new(network), networkClass},
		"fatal error": {
			func() (err error) {
				defer func() {
					if r := recover(); r != nil {
						err = &recovered{r}
					}
				}()
				panic("at the Disco")
			}(), fatalClass},
		"repeatable error": {new(retriable), repeatableClass},
		"nil error":        {nil, Unknown},
		"unknown error":    {os.ErrInvalid, Unknown},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, classifier.Classify(test.err, Unknown))
			assert.Equal(t, Unknown, (Classifier)(nil).Classify(test.err, Unknown))
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
			false,
		},
		"nil":   {nil, true},
		"empty": {make(Classifier), true},
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
	fatalClass      = "fatal"
	filesystemClass = "fs"
	repeatableClass = "repeatable"
	timeoutClass    = "timeout"
)

type network struct{}

func (err *network) Error() string   { return "network error" }
func (err *network) Timeout() bool   { return true }
func (err *network) Temporary() bool { return true }

type recovered struct{ panic interface{} }

func (err *recovered) Error() string      { return "recovered after panic" }
func (err *recovered) Cause() interface{} { return err.panic }

type retriable struct{}

func (err *retriable) Error() string   { return "retriable action error" }
func (err *retriable) Retriable() bool { return true }
