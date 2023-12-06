package models

import (
	"github.com/jmoiron/sqlx"
)

type Region struct {
	Id     uint   `db:"id"`
	Nombre string `db:"nombre"`
	Numero uint   `db:"numero"`
}

func (r *Region) List(db *sqlx.DB) ([]Region, error) {
	regiones := []Region{}
	err := db.Select(&regiones, "SELECT * FROM region")
	if err != nil {
		return nil, err
	}
	return regiones, nil
}

func (r *Region) Insert(db *sqlx.DB) error {
	_, err := db.Exec("INSERT INTO region (nombre, numero) VALUES ($1, $2)", r.Nombre, r.Numero)
	if err != nil {
		return err
	}
	return nil
}

func (r *Region) Update(db *sqlx.DB) error {
	_, err := db.Exec("UPDATE region SET nombre = $1, numero = $2 WHERE id = $3", r.Nombre, r.Numero, r.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Region) Delete(db *sqlx.DB) error {
	_, err := db.Exec("DELETE FROM region WHERE id = $1", r.Id)
	if err != nil {
		return err
	}
	return nil
}
