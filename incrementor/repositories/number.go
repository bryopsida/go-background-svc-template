package repositories

import (
	"encoding/json"

	"github.com/dgraph-io/badger"
)

type Number struct {
	ID     string
	Number uint64
}

type INumberRepository interface {
	Save(number Number) error
	FindByID(id string) (*Number, error)
	DeleteByID(id string) error
}

type badgerNumberRepository struct {
	db *badger.DB
}

func NewBadgerNumberRepository(db *badger.DB) INumberRepository {
	return &badgerNumberRepository{db: db}
}

func (r *badgerNumberRepository) Save(number Number) error {
	return r.db.Update(func(txn *badger.Txn) error {
		data, err := json.Marshal(number)
		if err != nil {
			return err
		}
		return txn.Set([]byte(number.ID), data)
	})
}

func (r *badgerNumberRepository) FindByID(id string) (*Number, error) {
	var number Number
	err := r.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(id))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			return json.Unmarshal(val, &number)
		})
	})
	if err != nil {
		return nil, err
	}
	return &number, nil
}

func (r *badgerNumberRepository) DeleteByID(id string) error {
	return r.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(id))
	})
}
