package safe

import "github.com/pkg/errors"

// Do reliably runs the action and captures panic as its error.
//
//  go safe.Do(func() error { return unsafe.Work() }, func(err error) { log.Println(err) })
//
func Do(action func() error, handler func(error)) {
	var err error
	defer func() {
		if err != nil {
			handler(err)
		}
	}()
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("safe panic: %#+v", r)
		}
	}()
	err = action()
}
