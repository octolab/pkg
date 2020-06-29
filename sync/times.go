package sync

import (
	"sync"
	"sync/atomic"
)

// Times add possibility to Do some action the N times.
//
//  barrier := sync.Times(100)
//
//  func Call(action func() error) {
//  	barrier.Do(func () {
//  		if err := action(); err != nil {
//  			logger.WithError(err).Error("send the error to the Sentry only 100 times")
//  		}
//  	})
//  }
//
func Times(n uint32) interface{ Do(func()) } {
	return &times{limit: n}
}

type times struct {
	limit uint32
	mx    sync.Mutex
	done  uint32
}

func (times *times) Do(fn func()) {
	if atomic.LoadUint32(&times.done) >= times.limit {
		return
	}
	times.doSlow(fn)
}

func (times *times) doSlow(fn func()) {
	times.mx.Lock()
	defer times.mx.Unlock()
	if times.done < times.limit {
		fn()
		atomic.AddUint32(&times.done, 1) // store only success case
	}
}
