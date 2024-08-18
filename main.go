package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/bryopsida/go-background-svc-template/config"
	"github.com/bryopsida/go-background-svc-template/datastore"
	"github.com/bryopsida/go-background-svc-template/incrementor"
	"github.com/bryopsida/go-background-svc-template/incrementor/repositories"
)

func main() {
	slog.Info("Starting up...")

	config := config.NewViperConfig()
	slog.Info("Database path", "path", config.GetDatabasePath())

	db, err := datastore.GetDatabase(config)
	if err != nil {
		slog.Error("Error opening database", "error", err)
		panic(err)
	}
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())
	repo := repositories.NewBadgerNumberRepository(db)
	// run routine 1
	go incrementor.Increment(ctx, repo)
	// run routine 2
	go incrementor.Print(ctx, repo)

	// wait for signal, then cancel context and exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// blocking wait for signal, while other routines execute
	sig := <-sigChan
	slog.Info("Received signal", "signal", sig)
	cancel()
	slog.Info("Shutting down gracefully...")
}
