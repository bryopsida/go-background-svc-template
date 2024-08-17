package incrementor

import (
	"context"
	"log/slog"
	"time"

	"github.com/bryopsida/go-background-svc-template/incrementor/repositories"
)

func Print(ctx context.Context, repo *repositories.INumberRepository) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(1 * time.Second)
			number, err := (*repo).FindByID(id)
			if err != nil {
				slog.Error("Error finding record", "error", err)
			} else {
				slog.Info("Current number", "number", number.Number)
			}
		}
	}
}
