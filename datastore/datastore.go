package datastore

import (
	"os"
	"path"

	"github.com/bryopsida/go-background-svc-template/interfaces"
	"github.com/dgraph-io/badger"
)

func GetDatabase(config interfaces.IConfig) (*badger.DB, error) {
	dbPath := config.GetDatabasePath()
	dbDir := path.Dir(dbPath)
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		os.MkdirAll(dbDir, os.ModePerm)
	}
	opts := badger.DefaultOptions(config.GetDatabasePath())
	return badger.Open(opts)
}
