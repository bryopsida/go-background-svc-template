package datastore

import (
	"log/slog"
	"os"
	"path"

	"github.com/bryopsida/go-background-svc-template/interfaces"
	"github.com/dgraph-io/badger"
)

func GetDatabase(config interfaces.IConfig) (*badger.DB, error) {
	dbPath := config.GetDatabasePath()
	dbDir := path.Dir(dbPath)
	_, err := os.Stat(dbDir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dbDir, os.ModePerm)
		if err != nil {
			slog.Error("Error creating database directory", "error", err)
		}
	}
	opts := badger.DefaultOptions(config.GetDatabasePath())
	return badger.Open(opts)
}
