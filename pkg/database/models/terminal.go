package models

import (
	"github.com/jmoiron/sqlx"
)

type Terminal struct {
	Id        uint    `db:"id"`
	Nombre    string  `db:"nombre"`
	Longitud  float64 `db:"longitud"`
	Latitud   float64 `db:"latitud"`
	Direccion string  `db:"direccion"`
	IdCiudad  uint    `db:"id_ciudad"`
}

func (t *Terminal) List(db *sqlx.DB) ([]Terminal, error) {
	terminales := []Terminal{}
	err := db.Select(&terminales, "SELECT * FROM terminal")
	if err != nil {
		return nil, err
	}
	return terminales, nil
}

func (t *Terminal) Insert(db *sqlx.DB) error {
	result := db.QueryRow(`
		INSERT INTO terminal (nombre, longitud, latitud, direccion, id_ciudad)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`, t.Nombre, t.Longitud, t.Latitud, t.Direccion, t.IdCiudad)
	err := result.Scan(&t.Id)
	if err != nil {
		return err
	}
	return nil
}

func (t *Terminal) Update(db *sqlx.DB) error {
	_, err := db.Exec(`
		UPDATE terminal
		SET
			nombre = $1, longitud = $2, latitud = $3, direccion = $4, id_ciudad = $5
		WHERE id = $6
	`, t.Nombre, t.Longitud, t.Latitud, t.Direccion, t.IdCiudad, t.Id)
	if err != nil {
		return err
	}
	return nil
}

func (t *Terminal) Delete(db *sqlx.DB) error {
	_, err := db.Exec("DELETE FROM terminal WHERE id = $1", t.Id)
	if err != nil {
		return err
	}
	return nil
}
