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
	const (
		osErrorClass  = "os"
		netErrorClass = "network"
	)

	_ = new(custom)

	classifier := Classifier{}.
		ClassifyAs(os.ErrExist, osErrorClass).
		ClassifyAs(os.ErrNotExist, osErrorClass).
		ClassifyAs(new(net.AddrError), netErrorClass).
		ClassifyAs((net.Error)(nil), netErrorClass)

	tests := map[string]struct {
		err      error
		expected string
	}{
		"os naked error":   {os.ErrExist, osErrorClass},
		"os wrapped error": {errors.Wrap(os.ErrExist, "wrapped"), osErrorClass},
		//"network error":    {new(net.AddrError), netErrorClass},
		//"custom error": {new(custom), netErrorClass},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, classifier.Classify(test.err, Unknown))
		})
	}
}

// helper

type custom struct{}

func (err *custom) Error() string   { return "custom" }
func (err *custom) Timeout() bool   { return true }
func (err *custom) Temporary() bool { return true }
