package safe

import "sync"

// A Job provides a top level API above the sync.WaitGroup.
//
//  job := new(safe.Job)
//  for _, action := range jobs {
//  	job.Do(action, logger)
//  }
//  job.Wait()
//
// It looks a little bit similar to the golang.org/x/sync/errgroup.
type Job sync.WaitGroup

// Do calls the given function in a new goroutine.
// If an error is not nil, it passes it to the handler.
func (job *Job) Do(action func() error, handler func(error)) {
	(*sync.WaitGroup)(job).Add(1)
	go Do(func() error {
		if err := action(); err != nil {
			return err
		}
		(*sync.WaitGroup)(job).Done()
		return nil
	}, func(err error) {
		handler(err)
		(*sync.WaitGroup)(job).Done()
	})
}

// Wait blocks until the sync.WaitGroup counter is zero.
func (job *Job) Wait() {
	(*sync.WaitGroup)(job).Wait()
}
