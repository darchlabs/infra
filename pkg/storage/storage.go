package storage

import (
	"github.com/jmoiron/sqlx"
)

type Store struct {
	DB *sqlx.DB
}

// NewPostgres ...
func NewPostgres(dsn string) (*Store, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &Store{
		DB: db,
	}, nil
}
