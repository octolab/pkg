package timer_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/fn/timer"
)

func TestStopwatch(t *testing.T) {
	var compare time.Duration

	duration := Stopwatch(func() {
		start := time.Now()
		time.Sleep(time.Millisecond)
		compare = time.Since(start)
	})

	assert.True(t, compare < duration)
}
