package main

import (
	"context"
	"os"

	"github.com/hrz8/silver-bassoon/pkg/logger"
)

func main() {
	ctx := context.Background()
	db := connect(ctx)
	err := start(db)

	defer func() {
		db.Close(ctx)
		logger.Debug("cleaning up...")

		os.Exit(0)
	}()

	select {
	case <-waitShutdown():
		logger.Info("shutting down...")
	case e := <-err:
		if e != nil {
			logger.Fatal("cannot start the app", e)
		}
	}
}
