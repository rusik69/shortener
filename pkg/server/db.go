package server

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
)

// DB represents a BadgerDB instance
type DB struct {
	db *badger.DB
}

// NewDB creates a new BadgerDB instance
func NewDB(path string) (*DB, error) {
	opts := badger.DefaultOptions(path)
	db, err := badger.Open(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	return &DB{db: db}, nil
}

// Set stores a key-value pair in the database
func (d *DB) Set(key, value string) error {
	return d.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(value))
	})
}

// Get retrieves a value from the database by key
func (d *DB) Get(key string) (string, error) {
	var value string
	err := d.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			value = string(val)
			return nil
		})
	})
	if err != nil {
		return "", err
	}
	return value, nil
}

// Close closes the database
func (d *DB) Close() error {
	return d.db.Close()
}
