// +build go1.13

package errors

import "errors"

// Unknown can be used as fallback class name of error classification.
const Unknown = "unknown"

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
			if errors.Is(err, target) {
				return class
			}
		}
	}
	return fallback
}

// ClassifyAs unwraps the error and store it as a class inside.
//
//  const network = "network"
//  global.Classifier.
//  	ClassifyAs(new(net.AddrError), network).
//  	ClassifyAs(new(net.DNSError), network).
//  	ClassifyAs(new(net.InvalidAddrError), network).
//  	ClassifyAs((net.Error)(nil), network)
//
func (classifier Classifier) ClassifyAs(err error, class string) Classifier {
	if err = Unwrap(err); err == nil {
		return classifier
	}
	for _, target := range classifier[class] {
		if errors.Is(err, target) {
			return classifier
		}
	}
	classifier[class] = append(classifier[class], err)
	return classifier
}
