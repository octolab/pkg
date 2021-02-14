package testing

import (
	"context"
	"time"
)

func WithDeadline(parent context.Context, src interface{}, delta time.Duration) (context.Context, context.CancelFunc) {
	deadline, is := src.(interface {
		Deadline() (time.Time, bool)
	})
	if !is {
		return context.WithCancel(parent)
	}
	value, has := deadline.Deadline()
	if !has {
		return context.WithCancel(parent)
	}
	return context.WithDeadline(parent, value.Add(-delta))
}
