package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hrz8/silver-bassoon/pkg/logger"
)

func waitShutdown() <-chan struct{} {
	wait := make(chan struct{})

	go func() {
		exit := make(chan os.Signal, 1)

		signal.Notify(
			exit,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGHUP,
			syscall.SIGQUIT,
		)

		sig := <-exit

		logger.Warn(fmt.Sprintf("%s signal triggered", sig))

		close(wait)
	}()

	return wait
}
