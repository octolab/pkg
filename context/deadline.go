package context

import (
	"context"
	"time"
)

type Deadliner interface {
	Deadline() (time.Time, bool)
}

func WithDeadline(parent context.Context, src Deadliner, delta time.Duration) (context.Context, context.CancelFunc) {
	deadline, has := src.Deadline()
	if !has {
		return context.WithCancel(parent)
	}
	return context.WithDeadline(parent, deadline.Add(-delta))
}
