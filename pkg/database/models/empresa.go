package models

import (
	"github.com/jmoiron/sqlx"
)

type Empresa struct {
	Id     uint   `db:"id"`
	Nombre string `db:"nombre"`
}

func (e *Empresa) List(db *sqlx.DB) ([]Empresa, error) {
	empresas := []Empresa{}
	err := db.Select(&empresas, "SELECT * FROM empresa")
	if err != nil {
		return nil, err
	}
	return empresas, nil
}

func (e *Empresa) Insert(db *sqlx.DB) error {
	result := db.QueryRow("INSERT INTO empresa (nombre) VALUES ($1) RETURNING id", e.Nombre)
	err := result.Scan(&e.Id)
	if err != nil {
		return err
	}
	return nil
}

func (e *Empresa) Update(db *sqlx.DB) error {
	_, err := db.Exec("UPDATE empresa SET nombre = $1 WHERE id = $2", e.Nombre, e.Id)
	if err != nil {
		return err
	}
	return nil
}

func (e *Empresa) Delete(db *sqlx.DB) error {
	_, err := db.Exec("DELETE FROM empresa WHERE id = $1", e.Id)
	if err != nil {
		return err
	}
	return nil
}
