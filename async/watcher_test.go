package async_test

import (
	"context"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/async"
	"go.octolab.org/sequence"
)

func TestWatcher(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), 200*time.Millisecond)
		defer cancel()

		var spy uint64

		r := rand.New(rand.NewSource(1))
		watchdog := Watcher(ctx, 50*time.Millisecond)
		watchdog.Watch("incrementer", func() { atomic.AddUint64(&spy, 1) })
		watchdog.Watch("duplicate", func() { atomic.AddUint64(&spy, 1) })
		for range sequence.Simple(1 + r.Intn(5)) {
			watchdog.Start()
		}
		watchdog.Forget("duplicate")

		<-ctx.Done()
		for range sequence.Simple(1 + r.Intn(5)) {
			watchdog.Stop()
		}
		assert.True(t, 0 < spy && spy <= 4, spy)
	})

	t.Run("failure", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), 200*time.Millisecond)
		defer cancel()

		var spy uint64

		watchdog := Watcher(ctx, 50*time.Millisecond)
		watchdog.Watch("incrementer", func() { atomic.AddUint64(&spy, 1) })
		watchdog.Watch("breaker", func() { panic("fail") })
		watchdog.Start()

		<-ctx.Done()
		watchdog.Stop()
		assert.True(t, 0 <= spy && spy <= 1, spy)
	})
}
