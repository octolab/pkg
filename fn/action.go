package fn

import "context"

func HasNoError(action func()) func() error {
	return func() error {
		action()
		return nil
	}
}

func HoldContext(ctx context.Context, action func(context.Context) error) func() error {
	return func() error { return action(ctx) }
}
