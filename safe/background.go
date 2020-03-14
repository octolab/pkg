package safe

import "sync"

type Job interface {
	Do(func() error, func(error))
	Wait()
}

func Background() Job {
	return &background{}
}

type background sync.WaitGroup

func (bg *background) Do(action func() error, handler func(error)) {
	(*sync.WaitGroup)(bg).Add(1)
	go Do(func() error {
		defer (*sync.WaitGroup)(bg).Done()
		return action()
	}, handler)
}

func (bg *background) Wait() {
	(*sync.WaitGroup)(bg).Wait()
}
