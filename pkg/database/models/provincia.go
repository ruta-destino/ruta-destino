package models

import (
	"github.com/jmoiron/sqlx"
)

type Provincia struct {
	Id       uint   `db:"id"`
	Nombre   string `db:"nombre"`
	IdRegion uint   `db:"id_region"`
}

func (p *Provincia) List(db *sqlx.DB) ([]Provincia, error) {
	provincias := []Provincia{}
	err := db.Select(&provincias, "SELECT * FROM provincia")
	if err != nil {
		return nil, err
	}
	return provincias, nil
}

func (p *Provincia) Insert(db *sqlx.DB) error {
	_, err := db.Exec("INSERT INTO provincia (nombre, id_region) VALUES ($1, $2)", p.Nombre, p.IdRegion)
	if err != nil {
		return err
	}
	return nil
}

func (p *Provincia) Update(db *sqlx.DB) error {
	_, err := db.Exec("UPDATE provincia SET nombre = $1, id_region = $2 WHERE id = $3", p.Nombre, p.IdRegion, p.Id)
	if err != nil {
		return err
	}
	return nil
}

func (p *Provincia) Delete(db *sqlx.DB) error {
	_, err := db.Exec("DELETE FROM provincia WHERE id = $1", p.Id)
	if err != nil {
		return err
	}
	return nil
}
