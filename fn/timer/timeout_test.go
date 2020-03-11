package timer_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/fn/timer"
)

func TestDeadline(t *testing.T) {
	type tuple struct {
		deadline time.Time
		present  bool
	}

	now, delta := time.Now(), time.Millisecond

	tests := map[string]struct {
		correction float32
		threshold  time.Duration
		payload    tuple
		expected   time.Time
	}{
		"fallback": {
			0.95,
			time.Millisecond,
			tuple{now.Add(time.Second), false},
			now.Add(time.Second),
		},
		"threshold": {
			0.95,
			time.Millisecond,
			tuple{now.Add(time.Second), true},
			now.Add(time.Second - time.Millisecond),
		},
		"corrected": {
			0.95,
			100 * time.Millisecond,
			tuple{now.Add(time.Second), true},
			now.Add(time.Second - 50*time.Millisecond),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			corrector := Deadline(test.correction, test.threshold)
			deadline, present := test.payload.deadline, test.payload.present
			assert.WithinDuration(t, test.expected, corrector(deadline, present), delta)
		})
	}
}

func TestTimeout(t *testing.T) {
	type tuple struct {
		timeout time.Duration
		present bool
	}

	tests := map[string]struct {
		correction float32
		threshold  time.Duration
		payload    tuple
		expected   time.Duration
	}{
		"fallback": {
			0.95,
			time.Millisecond,
			tuple{time.Second, false},
			time.Second,
		},
		"threshold": {
			0.95,
			time.Millisecond,
			tuple{time.Second, true},
			time.Second - time.Millisecond,
		},
		"corrected": {
			0.95,
			100 * time.Millisecond,
			tuple{time.Second, true},
			time.Second - 50*time.Millisecond,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			corrector := Timeout(test.correction, test.threshold)
			timeout, present := test.payload.timeout, test.payload.present
			assert.Equal(t, test.expected, corrector(timeout, present))
		})
	}
}
