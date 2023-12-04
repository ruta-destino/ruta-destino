package models

import (
	"github.com/jmoiron/sqlx"
)

type Region struct {
	Id     uint   `db:"id"`
	Nombre string `db:"nombre"`
}

func (r *Region) List(db *sqlx.DB) []Region {
	regiones := []Region{}
	db.Select(&regiones, "SELECT * FROM region")
	return regiones
}

func (r *Region) Insert(db *sqlx.DB) error {
	_, err := db.Exec("INSERT INTO region (nombre) VALUES ($1)", r.Nombre)
	if err != nil {
		return err
	}
	return nil
}
