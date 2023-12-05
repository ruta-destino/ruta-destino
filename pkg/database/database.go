package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB

// Opens a new connection to the database and stores it as a global package
// variable. Works as a singleton.
func Open(connectionString string) (*sqlx.DB, error) {
	if Db != nil {
		fmt.Println("Reused connection")
		return Db, nil
	}
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	Db = db
	return Db, nil
}
