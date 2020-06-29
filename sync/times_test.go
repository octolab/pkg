package sync_test

import (
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.octolab.org/sequence"
	. "go.octolab.org/sync"
)

func TestTimes(t *testing.T) {
	var counter int32

	limiter := Times(100)
	for range sequence.Simple(1000) {
		limiter.Do(func() { atomic.AddInt32(&counter, 1) })
	}

	assert.Equal(t, int32(100), counter)
}
