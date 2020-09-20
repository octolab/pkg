package async

import (
	"context"
	"sync"
	"time"

	"go.octolab.org/safe"
)

func New(ctx context.Context, frequency time.Duration) *watcher {
	ticker := time.NewTicker(frequency)
	return &watcher{ctx: ctx, ticker: ticker, stuff: make(map[string]func())}
}

type watcher struct {
	ctx         context.Context
	start, stop sync.Once
	ticker      *time.Ticker
	guard       sync.RWMutex
	stuff       map[string]func()
}

func (watcher *watcher) Watch(name string, callback func()) {
	watcher.guard.Lock()
	watcher.stuff[name] = callback
	watcher.guard.Unlock()
}

func (watcher *watcher) Forget(name string) {
	watcher.guard.Lock()
	delete(watcher.stuff, name)
	watcher.guard.Unlock()
}

func (watcher *watcher) Start() {
	watcher.start.Do(func() {
		go safe.Do(func() error {
			for {
				select {
				case <-watcher.ctx.Done():
					watcher.Stop()
					return nil
				case <-watcher.ticker.C:
					watcher.guard.RLock()
					for _, fn := range watcher.stuff {
						fn()
					}
					watcher.guard.RUnlock()
				}
			}
		}, func(err error) {
			watcher.guard.RUnlock()
			watcher.Stop()
		})
	})
}

func (watcher *watcher) Stop() {
	watcher.stop.Do(func() { watcher.ticker.Stop() })
}
