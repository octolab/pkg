// +build go1.13

package errors_test

import (
	"net"
	"os"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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
		"naked error":      {os.ErrExist, filesystemClass},
		"wrapped error":    {errors.Wrap(os.ErrExist, "wrapped"), filesystemClass},
		"network error":    {new(net.AddrError), networkClass},
		"custom error":     {new(network), networkClass},
		"repeatable error": {new(retriable), repeatableClass},
		"nil error":        {nil, Unknown},
		"unknown error":    {os.ErrInvalid, Unknown},
		"fatal error": {
			func() (err error) {
				defer func() {
					if r := recover(); r != nil {
						err = &recovered{r}
					}
				}()
				panic("at the Disco")
			}(), fatalClass},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			require.True(t, classifier.Consistent())
			assert.Equal(t, test.expected, classifier.Classify(test.err, Unknown))
			assert.Equal(t, Unknown, (Classifier)(nil).Classify(test.err, Unknown))
		})
	}
}

func TestClassifier_ClassifyAs(t *testing.T) {
	tests := map[string]struct {
		classifier Classifier
		list       []error
	}{
		networkClass: {
			make(Classifier).ClassifyAs(networkClass, new(NetworkError)),
			[]error{
				new(net.AddrError),
				&net.DNSConfigError{Err: new(net.DNSError)},
				new(net.DNSError),
				new(net.InvalidAddrError),
				&net.OpError{Err: new(net.InvalidAddrError)},
				new(net.UnknownNetworkError),
			},
		},
		fatalClass: {
			make(Classifier).ClassifyAs(fatalClass, new(RecoveredError)),
			[]error{new(recovered)},
		},
		filesystemClass: {
			make(Classifier).ClassifyAs(filesystemClass,
				os.ErrClosed, os.ErrExist, os.ErrInvalid,
				os.ErrNoDeadline, os.ErrNotExist, os.ErrPermission),
			[]error{
				os.ErrClosed,
				os.ErrExist,
				os.ErrInvalid,
				os.ErrNoDeadline,
				os.ErrNotExist,
				os.ErrPermission,
			},
		},
		repeatableClass: {
			make(Classifier).ClassifyAs(repeatableClass, new(RetriableError)),
			[]error{
				new(retriable),
			},
		},
		temporaryClass: {
			make(Classifier).ClassifyAs(temporaryClass, new(TemporaryError)),
			[]error{
				new(net.AddrError),
				&net.DNSConfigError{Err: new(net.DNSError)},
				new(net.DNSError),
				new(net.InvalidAddrError),
				&net.OpError{Err: new(net.InvalidAddrError)},
				new(net.UnknownNetworkError),
			},
		},
		timeoutClass: {
			make(Classifier).ClassifyAs(timeoutClass, new(TimeoutError)),
			[]error{
				new(net.AddrError),
				&net.DNSConfigError{Err: new(net.DNSError)},
				new(net.DNSError),
				new(net.InvalidAddrError),
				&net.OpError{Err: new(net.InvalidAddrError)},
				new(net.UnknownNetworkError),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			require.True(t, test.classifier.Consistent())
			for _, err := range test.list {
				assert.NotEqual(t, Unknown, test.classifier.Classify(err, Unknown))
			}
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

func TestClassifier_Presets(t *testing.T) {
	t.Run(networkClass, func(t *testing.T) {
		err := new(NetworkError)
		assert.EqualError(t, err, "network error")
		assert.False(t, err.Temporary())
		assert.False(t, err.Timeout())
	})
	t.Run(fatalClass, func(t *testing.T) {
		err := new(RecoveredError)
		assert.EqualError(t, err, "recovered after panic")
		assert.Nil(t, err.Cause())
	})
	t.Run(repeatableClass, func(t *testing.T) {
		err := new(RetriableError)
		assert.EqualError(t, err, "retriable action error")
		assert.False(t, err.Retriable())
	})
	t.Run(temporaryClass, func(t *testing.T) {
		err := new(TemporaryError)
		assert.EqualError(t, err, "temporary error")
		assert.False(t, err.Temporary())
	})
	t.Run(timeoutClass, func(t *testing.T) {
		err := new(TimeoutError)
		assert.EqualError(t, err, "timeout error")
		assert.False(t, err.Timeout())
	})
}

// helper

const (
	networkClass    = "network"
	fatalClass      = "fatal"
	filesystemClass = "fs"
	repeatableClass = "repeatable"
	temporaryClass  = "temporary"
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
