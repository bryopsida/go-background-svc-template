package incrementor

import (
	"context"
	"log/slog"
	"strings"
	"time"

	"github.com/bryopsida/go-background-svc-template/incrementor/repositories"
)

const id = "incrementor"

func saveRecord(repo repositories.INumberRepository, number repositories.Number) {
	err := repo.Save(number)
	if err != nil {
		slog.Error("Error saving record", "error", err)
	}
}

func initializeRecord(repo repositories.INumberRepository) {
	number := repositories.Number{
		ID:     id,
		Number: 0,
	}
	saveRecord(repo, number)
}

func Increment(ctx context.Context, repo *repositories.INumberRepository) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(1 * time.Second)
			number, err := (*repo).FindByID(id)
			if err != nil {
				if strings.Contains(strings.ToLower(err.Error()), "key not found") {
					initializeRecord(*repo)
				} else {
					slog.Error("Error finding record", "error", err)
				}
			} else {
				number.Number++
				saveRecord(*repo, *number)
			}
		}
	}
}
