package services

import (
	"ruta-destino/pkg/database/models"

	"github.com/jmoiron/sqlx"
)

type Ciudad struct {
	Db *sqlx.DB
}

func NewCiudadService(db *sqlx.DB) *Ciudad {
	return &Ciudad{Db: db}
}

func (s *Ciudad) List() ([]models.Ciudad, error) {
	ciudades := []models.Ciudad{}
	err := s.Db.Select(&ciudades, `
		SELECT ciudad.*, provincia.nombre AS "nombre_provincia"
		FROM ciudad
		INNER JOIN provincia
		ON ciudad.id_provincia = provincia.id
	`)
	if err != nil {
		return nil, err
	}
	return ciudades, nil
}

func (s *Ciudad) Get(ciudadId uint) (*models.Ciudad, error) {
	ciudad := models.Ciudad{}
	err := s.Db.Get(&ciudad, `
		SELECT *
		FROM ciudad
		WHERE id = $1
	`, ciudadId)
	if err != nil {
		return nil, err
	}
	return &ciudad, nil
}

func (s *Ciudad) Insert(ciudad *models.Ciudad) error {
	result := s.Db.QueryRow(`
		INSERT INTO ciudad (nombre, id_provincia)
		VALUES ($1, $2)
		RETURNING id
	`, ciudad.Nombre, ciudad.IdProvincia)
	err := result.Scan(&ciudad.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Ciudad) Update(idCiudad uint, ciudad *models.Ciudad) error {
	_, err := s.Db.Exec(`
		UPDATE ciudad
		SET nombre = $1, id_provincia = $2
		WHERE id = $3
	`, ciudad.Nombre, ciudad.IdProvincia, idCiudad)
	if err != nil {
		return err
	}
	return nil
}

func (s *Ciudad) Delete(idCiudad uint) error {
	_, err := s.Db.Exec(`
		DELETE FROM ciudad
		WHERE id = $1
	`, idCiudad)
	if err != nil {
		return err
	}
	return nil
}
