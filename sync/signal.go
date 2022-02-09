package sync

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.octolab.org/errors"
)

// ErrSignalTrapped is returned by the SignalTrap.Wait
// when the expected signals caught.
const ErrSignalTrapped errors.Message = "signal trapped"

// Termination returns trap for termination signals.
//
//	server := new(http.Server)
//	go safe.Do(server.ListenAndServe, func(err error) { log.Println(err) })
//
//	err := sync.Termination().Wait(context.Background())
//	if err == sync.ErrSignalTrapped {
//		log.Println("shutting down the server", server.Shutdown(context.Background()))
//	}
func Termination() SignalTrap {
	trap := make(chan os.Signal, 3)
	signal.Notify(trap, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	return trap
}

// SignalTrap wraps os.Signal channel to provide high level API above it.
//
//	trap := make(chan os.Signal)
//	signal.Notify(trap, os.Interrupt)
//	SignalTrap(trap).Wait(context.Background())
type SignalTrap chan os.Signal

// Wait blocks until one of the expected signals caught
// or the Context closed. It unregisters from the notification
// and closes itself.
func (trap SignalTrap) Wait(ctx context.Context) error {
	defer close(trap)
	defer signal.Stop(trap)

	select {
	case <-trap:
		return ErrSignalTrapped
	case <-ctx.Done():
		return ctx.Err()
	}
}
