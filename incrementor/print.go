package incrementor

import (
	"context"
	"log/slog"
	"time"

	"github.com/bryopsida/go-background-svc-template/interfaces"
)

// Print prints the current number
// - ctx: context.Context to signal the function to stop
// - repo: interfaces.INumberRepository to interact with the database
func Print(ctx context.Context, repo interfaces.INumberRepository) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(1 * time.Second)
			number, err := repo.FindByID(getID())
			if err != nil {
				slog.Error("Error finding record", "error", err)
			} else {
				slog.Info("Current number", "number", number.Number)
			}
		}
	}
}
