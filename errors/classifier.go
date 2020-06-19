// +build go1.13

package errors

import (
	"errors"
	"net"
)

// Classifier provides functionality to classify errors
// and represents them as a string, e.g. for metrics system.
type Classifier map[string][]error

// Classify classifies the error and returns its class name.
//
//  func (service *Service) Do(ctx context.Context, payload interface{}) {
//  	resp, err := service.proxy.Call(ctx, Data{Payload: payload})
//  	if err != nil {
//  		service.telemetry.Increment(global.Classifier.Classify(err, errors.Unknown))
//  		...
//  	}
//  	...
//  }
//
func (classifier Classifier) Classify(err error, fallback string) string {
	if err = Unwrap(err); err == nil {
		return fallback
	}
	for class, list := range classifier {
		for _, target := range list {
			if errors.Is(target, err) {
				return class
			}
		}
	}
	return fallback
}

// ClassifyAs unwraps the errors and stores them with the class name.
//
//  classifier := make(errors.Classifier).
//  	ClassifyAs("network", new(errors.NetworkError)).
//  	ClassifyAs("fs", os.ErrExist, os.ErrNotExist)
//
func (classifier Classifier) ClassifyAs(class string, list ...error) Classifier {
	for _, err := range list {
		if err = Unwrap(err); err == nil {
			continue
		}
		var present bool
		for _, target := range classifier[class] {
			if errors.Is(target, err) {
				present = true
				break
			}
		}
		if !present {
			classifier[class] = append(classifier[class], err)
		}
	}
	return classifier
}

// Consistent checks that different groups don't contain similar errors.
func (classifier Classifier) Consistent() bool {
	return true
}

// Unknown can be used as the fallback class name of error classification.
const Unknown = "unknown"

// NetworkError can check network errors.
type NetworkError struct{}

func (*NetworkError) Error() string   { return "network error" }
func (*NetworkError) Temporary() bool { return false }
func (*NetworkError) Timeout() bool   { return false }

// Is reports whether the error matches network error class.
func (*NetworkError) Is(err error) bool {
	_, is := err.(net.Error)
	return is
}

// TemporaryError can check temporary errors.
type TemporaryError struct{}

func (*TemporaryError) Error() string   { return "temporary error" }
func (*TemporaryError) Temporary() bool { return false }

// Is reports whether the error matches temporary error class.
func (*TemporaryError) Is(err error) bool {
	_, is := err.(interface{ Temporary() bool })
	return is
}

// TimeoutError can check timeout errors.
type TimeoutError struct{}

func (*TimeoutError) Error() string { return "timeout error" }
func (*TimeoutError) Timeout() bool { return false }

// Is reports whether the error matches timeout error class.
func (*TimeoutError) Is(err error) bool {
	_, is := err.(interface{ Timeout() bool })
	return is
}
