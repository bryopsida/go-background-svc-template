package incrementor

import (
	"context"
	"log/slog"
	"strings"
	"time"

	"github.com/bryopsida/go-background-svc-template/interfaces"
)

func getID() string {
	return "incrementor"
}

func saveRecord(repo interfaces.INumberRepository, number interfaces.Number) {
	err := repo.Save(number)
	if err != nil {
		slog.Error("Error saving record", "error", err)
	}
}

func initializeRecord(repo interfaces.INumberRepository) {
	number := interfaces.Number{
		ID:     getID(),
		Number: 0,
	}
	saveRecord(repo, number)
}

// Increment increments the number in the record
// - ctx: context.Context to signal the function to stop
// - repo: interfaces.INumberRepository to interact with the database
func Increment(ctx context.Context, repo interfaces.INumberRepository) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(1 * time.Second)
			number, err := repo.FindByID(getID())
			if err != nil {
				if strings.Contains(strings.ToLower(err.Error()), "key not found") {
					initializeRecord(repo)
				} else {
					slog.Error("Error finding record", "error", err)
				}
			} else if number != nil {
				number.Number++
				saveRecord(repo, *number)
			}
		}
	}
}
