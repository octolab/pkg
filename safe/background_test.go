package safe_test

import (
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/safe"
)

func TestBackground(t *testing.T) {
	var spy uint32

	job := Background()
	for range make([]struct{}, 10) {
		job.Do(func() error {
			atomic.AddUint32(&spy, 1)
			return nil
		}, nil)
	}

	job.Wait()
	assert.Equal(t, uint32(10), spy)
}
