package sync_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/sync"
)

const delta = 10 * time.Millisecond

func TestTermination(t *testing.T) {
	tests := map[string]struct {
		breaker  func(cancel context.CancelFunc)
		expected error
	}{
		"break by signal": {
			func(cancel context.CancelFunc) {
				proc, err := os.FindProcess(os.Getpid())
				assert.NoError(t, err)
				assert.NoError(t, proc.Signal(os.Interrupt))
				time.AfterFunc(delta, cancel)
			},
			ErrSignalTrapped,
		},
		"break by context": {
			func(cancel context.CancelFunc) {
				cancel()
			},
			context.Canceled,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			trap := Termination()
			go test.breaker(cancel)
			assert.Equal(t, test.expected, trap.Wait(ctx))
		})
	}
}
