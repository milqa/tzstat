package grsdown

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	exitCodeOk       = 0
	exitCodeWatchdog = 1
)

const (
	shutdownTimeout = time.Second * 5
	watchdogTimeout = shutdownTimeout + time.Second*5
)

// Run f until interrupt.
func Run(ctx context.Context,
	f func(ctx context.Context) error,
) {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error {
		if err := f(ctx); err != nil {
			return err
		}
		return nil
	})

	go func() {
		// Guaranteed way to kill application.
		<-ctx.Done()

		// Context is canceled, giving application time to shut down gracefully.
		log.Print("Waiting for application shutdown")
		time.Sleep(watchdogTimeout)

		// Probably deadlock, forcing shutdown.
		log.Print("Graceful shutdown watchdog triggered: forcing shutdown")
		os.Exit(exitCodeWatchdog)
	}()

	if err := wg.Wait(); err != nil {
		if errors.Is(err, context.Canceled) {
			log.Print("Graceful shutdown")
			return
		}

		log.Fatal("Failed", err)
	}

	os.Exit(exitCodeOk)
}
