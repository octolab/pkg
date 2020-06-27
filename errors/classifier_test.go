// +build go1.13

package errors_test

import (
	"fmt"
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
		ClassifyAs(fatalClass, new(RecoveredError)).
		ClassifyAs(filesystemClass, os.ErrExist, os.ErrNotExist).
		ClassifyAs(networkClass, new(NetworkError)).
		ClassifyAs(repeatableClass, new(RetriableError)).
		ClassifyAs(Unknown, nil)

	tests := map[string]struct {
		err      error
		expected string
	}{
		"naked error":          {os.ErrExist, filesystemClass},
		"wrapped by pkg error": {errors.Wrap(os.ErrExist, "wrapped"), filesystemClass},
		"wrapped by fmt error": {fmt.Errorf("wrapped: %w", os.ErrExist), filesystemClass},
		"network error":        {new(net.AddrError), networkClass},
		"custom error":         {new(network), networkClass},
		"repeatable error":     {new(retriable), repeatableClass},
		"fatal error":          {fatal(), fatalClass},
		"nil error":            {nil, Unknown},
		"unknown error":        {os.ErrInvalid, Unknown},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			require.True(t, classifier.Consistent())
			assert.Equal(t, test.expected, classifier.Classify(test.err))
			assert.Equal(t, Unknown, (Classifier)(nil).Classify(test.err))
		})
	}
}

func TestClassifier_ClassifyAs(t *testing.T) {
	tests := map[string]struct {
		classifier Classifier
		list       []error
	}{
		fatalClass: {
			make(Classifier).ClassifyAs(fatalClass, new(RecoveredError)),
			[]error{
				new(recovered),
				fatal(),
			},
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
		repeatableClass: {
			make(Classifier).ClassifyAs(repeatableClass, new(RetriableError)),
			[]error{
				new(retriable),
			},
		},
		specificClass: {
			make(Classifier).ClassifyAs(specificClass, MessageError{Message: "dial tcp"}),
			[]error{
				errors.New("dial tcp: i/o timeout"),
				errors.New(
					(&net.OpError{
						Op:  "dial",
						Net: "tcp",
						Err: fmt.Errorf("i/o timeout"),
					}).Error(),
				),
			},
		},
		temporaryClass: {
			make(Classifier).ClassifyAs(temporaryClass, new(TemporaryError)),
			[]error{
				temporary(),
				&net.DNSConfigError{Err: temporary()},
				&net.OpError{Err: temporary()},
			},
		},
		timeoutClass: {
			make(Classifier).ClassifyAs(timeoutClass, new(TimeoutError)),
			[]error{
				timeout(),
				&net.DNSConfigError{Err: timeout()},
				&net.OpError{Err: timeout()},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			require.True(t, test.classifier.Consistent())
			for _, err := range test.list {
				assert.NotEqual(t, Unknown, test.classifier.Classify(err))
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
				ClassifyAs(filesystemClass, os.ErrExist, os.ErrNotExist).
				ClassifyAs(networkClass,
					new(TemporaryError), new(NetworkError), new(TimeoutError)),
			true,
		},
		"consistent subset case": {
			make(Classifier).
				ClassifyAs(networkClass,
					new(NetworkError), new(TemporaryError), new(TimeoutError)),
			true,
		},
		"not consistent": {
			make(Classifier).
				ClassifyAs(networkClass, new(NetworkError)).
				ClassifyAs(temporaryClass, new(TemporaryError)).
				ClassifyAs(timeoutClass, new(TimeoutError)),
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

func TestClassifier_Merge(t *testing.T) {
	classifierX := make(Classifier).ClassifyAs(fatalClass, new(RecoveredError))
	assert.True(t, classifierX.Consistent())
	assert.Equal(t, fatalClass, classifierX.Classify(new(recovered)))
	assert.Equal(t, Unknown, classifierX.Classify(new(network)))

	classifierY := make(Classifier).ClassifyAs(networkClass, new(NetworkError))
	assert.True(t, classifierY.Consistent())
	assert.Equal(t, networkClass, classifierY.Classify(new(network)))
	assert.Equal(t, Unknown, classifierY.Classify(new(recovered)))

	classifierZ := classifierY.Merge(classifierX)
	assert.True(t, classifierZ.Consistent())
	assert.Equal(t, fatalClass, classifierZ.Classify(new(recovered)))
	assert.Equal(t, networkClass, classifierZ.Classify(new(network)))
}

func TestClassifier_Presets(t *testing.T) {
	t.Run(fatalClass, func(t *testing.T) {
		err := new(RecoveredError)
		assert.EqualError(t, err, "recovered after panic")
		assert.Nil(t, err.Cause())
	})
	t.Run(networkClass, func(t *testing.T) {
		err := new(NetworkError)
		assert.EqualError(t, err, "network error")
		assert.True(t, err.Temporary())
		assert.True(t, err.Timeout())
	})
	t.Run(repeatableClass, func(t *testing.T) {
		err := new(RetriableError)
		assert.EqualError(t, err, "retriable action error")
		assert.True(t, err.Retriable())
	})
	t.Run(specificClass, func(t *testing.T) {
		err := MessageError{Message: "error"}
		assert.EqualError(t, err, err.Message)
	})
	t.Run(temporaryClass, func(t *testing.T) {
		err := new(TemporaryError)
		assert.EqualError(t, err, "temporary error")
		assert.True(t, err.Temporary())
	})
	t.Run(timeoutClass, func(t *testing.T) {
		err := new(TimeoutError)
		assert.EqualError(t, err, "timeout error")
		assert.True(t, err.Timeout())
	})
}

// helpers

const (
	fatalClass      = "fatal"
	filesystemClass = "fs"
	networkClass    = "network"
	repeatableClass = "repeatable"
	specificClass   = "specific"
	temporaryClass  = "temporary"
	timeoutClass    = "timeout"
)

func fatal() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = &recovered{r}
		}
	}()
	panic("at the Disco")
}

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

func temporary() error {
	err := new(net.DNSError)
	err.IsTemporary = true
	return err
}

func timeout() error {
	err := new(net.DNSError)
	err.IsTimeout = true
	return err
}
