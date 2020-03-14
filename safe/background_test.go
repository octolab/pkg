package safe_test

import (
	"errors"
	"sync"
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
		func() error { panic(errors.New("at the Disco")) },
		func(error) { atomic.AddUint32(&spy, 1) },
	)

	job.Wait()
	assert.Equal(t, uint32(12), spy)
}

// go test -run=NONE -bench=BenchmarkBackground ./safe
// 1293283               917 ns/op              80 B/op          3 allocs/op
// 1260637               957 ns/op              96 B/op          4 allocs/op
//  125277              8455 ns/op             504 B/op         10 allocs/op
//  677760              1520 ns/op              32 B/op          2 allocs/op
func BenchmarkBackground(b *testing.B) {
	b.Run("without error", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			job := Background()
			job.Do(func() error { return nil }, nil)
			job.Wait()
		}
	})
	b.Run("with error", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			job := Background()
			job.Do(func() error { return errors.New("test") }, func(error) {})
			job.Wait()
		}
	})
	b.Run("with panic", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			job := Background()
			job.Do(func() error { panic(errors.New("at the Disco")) }, func(error) {})
			job.Wait()
		}
	})
	b.Run("wait group", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				defer func() { recover() }()
				panic(errors.New("at the Disco"))
			}()
			wg.Wait()
		}
	})
}
