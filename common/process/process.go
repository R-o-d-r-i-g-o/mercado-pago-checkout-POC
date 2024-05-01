package process

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type GracefulShutdownOptions struct {
	TimeoutSeconds time.Duration
	Quit           chan os.Signal
}

// GracefulShutdown is responsible for capture os interrupt signals to start graceful shutdown flow for more resilient shutdowns
func GracefulShutdown(shutdownCallback func(context.Context), opts ...GracefulShutdownOptions) {
	const EXIT_SIGNAL_CODE int = 1
	var options = GracefulShutdownOptions{
		TimeoutSeconds: 5,
		Quit:           make(chan os.Signal, EXIT_SIGNAL_CODE),
	}

	for _, opt := range opts {
		if opt.TimeoutSeconds != 0 {
			options.TimeoutSeconds = opt.TimeoutSeconds
		}

		if opt.Quit != nil {
			options.Quit = opt.Quit
		}
	}

	signal.Notify(options.Quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-options.Quit

	log.Println("gracefully shutdown process...")

	ctx, cancel := context.WithTimeout(context.Background(), options.TimeoutSeconds*time.Second)
	defer cancel()
	defer signal.Stop(options.Quit)

	go shutdownCallback(ctx)

	<-ctx.Done()

	log.Println("process exiting...")
}
