package database

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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

func ProcessPostgresError(err *pq.Error) error {
	message := ""
	switch err.Code {
	case "23505":
		pattern := `Key \(([^()]*)\)=\(([^()]*)\)`
		r, _ := regexp.Compile(pattern)
		match := r.FindStringSubmatch(err.Detail)
		columna, valor := match[1], match[2]
		message = fmt.Sprintf("%v(%v) ya existe", columna, valor)
	case "23503":
		pattern := `[^"]*"([^"]*)".*on table "([^"]*)"`
		r, _ := regexp.Compile(pattern)
		match := r.FindStringSubmatch(err.Message)
		tabla, tablaReferencia := match[1], match[2]
		message = fmt.Sprintf("%v está referenciada en %v", tabla, tablaReferencia)
	}
	if message == "" {
		fmt.Printf("%#v\n", err)
		return errors.New("FIXME")
	}
	return errors.New(message)
}
