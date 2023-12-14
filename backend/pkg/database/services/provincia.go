package services

import (
	"ruta-destino/pkg/database/models"

	"github.com/jmoiron/sqlx"
)

type Provincia struct {
	Db *sqlx.DB
}

func NewProvinciaService(db *sqlx.DB) *Provincia {
	return &Provincia{Db: db}
}

func (s *Provincia) List() ([]models.Provincia, error) {
	provincias := []models.Provincia{}
	err := s.Db.Select(&provincias, `
		SELECT provincia.*, region.nombre AS "nombre_region"
		FROM provincia
		INNER JOIN region
		ON provincia.id_region = region.id
	`)
	if err != nil {
		return nil, err
	}
	return provincias, nil
}

func (s *Provincia) Get(provinciaId uint) (*models.Provincia, error) {
	provincia := models.Provincia{}
	err := s.Db.Get(&provincia, `
		SELECT *
		FROM provincia
		WHERE id = $1
	`, provinciaId)
	if err != nil {
		return nil, err
	}
	return &provincia, nil
}

func (s *Provincia) Insert(provincia *models.Provincia) error {
	result := s.Db.QueryRow(`
		INSERT INTO provincia (nombre, id_region)
		VALUES ($1, $2)
		RETURNING id
	`, provincia.Nombre, provincia.IdRegion)
	err := result.Scan(&provincia.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Provincia) Update(idProvincia uint, provincia *models.Provincia) error {
	_, err := s.Db.Exec(`
		UPDATE provincia
		SET nombre = $1, id_region = $2
		WHERE id = $3
	`, provincia.Nombre, provincia.IdRegion, idProvincia)
	if err != nil {
		return err
	}
	return nil
}

func (s *Provincia) Delete(idProvincia uint) error {
	_, err := s.Db.Exec(`
		DELETE FROM provincia
		WHERE id = $1
	`, idProvincia)
	if err != nil {
		return err
	}
	return nil
}
