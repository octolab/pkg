package safe_test

import (
	"errors"
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
	job.Do(
		func() error { return errors.New("test") },
		func(error) { atomic.AddUint32(&spy, 1) },
	)
	job.Do(
		func() error { panic("at the Disco") },
		func(error) { atomic.AddUint32(&spy, 1) },
	)

	job.Wait()
	assert.Equal(t, uint32(12), spy)
}
