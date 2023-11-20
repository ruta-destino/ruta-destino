package database

import (
	"github.com/jmoiron/sqlx"
)

func Connect(connectionString string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
