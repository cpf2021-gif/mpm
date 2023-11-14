package model

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewSqlite() (*DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DB{DB: db}, nil
}

func MustNewSqlite() *DB {
	db, err := NewSqlite()
	if err != nil {
		fmt.Printf("open db error: %v", err)
		os.Exit(1)
	}

	return db
}
