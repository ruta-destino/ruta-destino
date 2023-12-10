package models

import (
	"github.com/jmoiron/sqlx"
)

type Ciudad struct {
	Id          uint   `db:"id"`
	Nombre      string `db:"nombre"`
	IdProvincia uint   `db:"id_provincia"`
}

func (c *Ciudad) List(db *sqlx.DB) ([]Ciudad, error) {
	ciudades := []Ciudad{}
	err := db.Select(&ciudades, `
		SELECT * FROM ciudad
	`)
	if err != nil {
		return nil, err
	}
	return ciudades, nil
}

func (c *Ciudad) Insert(db *sqlx.DB) error {
	result := db.QueryRow(`
		INSERT INTO ciudad (nombre, id_provincia)
		VALUES ($1, $2) RETURNING id
	`, c.Nombre, c.IdProvincia)
	err := result.Scan(&c.Id)
	if err != nil {
		return err
	}
	return nil
}

func (c *Ciudad) Update(db *sqlx.DB) error {
	_, err := db.Exec(`
		UPDATE ciudad
		SET nombre = $1, id_provincia = $2
		WHERE id = $3
	`, c.Nombre, c.IdProvincia, c.Id)
	if err != nil {
		return err
	}
	return nil
}

func (c *Ciudad) Delete(db *sqlx.DB) error {
	_, err := db.Exec(`
		DELETE FROM ciudad
		WHERE id = $1
	`, c.Id)
	if err != nil {
		return err
	}
	return nil
}
