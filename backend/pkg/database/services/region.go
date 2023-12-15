package services

import (
	"ruta-destino/pkg/database/models"

	"github.com/jmoiron/sqlx"
)

type Region struct {
	Db *sqlx.DB
}

func NewRegionService(db *sqlx.DB) *Region {
	return &Region{Db: db}
}

func (s *Region) List() ([]models.Region, error) {
	regiones := []models.Region{}
	err := s.Db.Select(&regiones, `
		SELECT *
		FROM region
		ORDER BY numero
	`)
	if err != nil {
		return nil, err
	}
	return regiones, nil
}

func (s *Region) Get(regionId uint) (*models.Region, error) {
	region := models.Region{}
	err := s.Db.Get(&region, `
		SELECT *
		FROM region
		WHERE id = $1
	`, regionId)
	if err != nil {
		return nil, err
	}
	return &region, nil
}

func (s *Region) Insert(region *models.Region) error {
	result := s.Db.QueryRow(`
		INSERT INTO region (nombre, numero)
		VALUES ($1, $2)
		RETURNING id
	`, region.Nombre, region.Numero)
	err := result.Scan(&region.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Region) Update(regionId uint, region *models.Region) error {
	_, err := s.Db.Exec(`
		UPDATE region
		SET nombre = $1, numero = $2
		WHERE id = $3
	`, region.Nombre, region.Numero, regionId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Region) Delete(regionId uint) error {
	_, err := s.Db.Exec(`
		DELETE FROM region
		WHERE id = $1
	`, regionId)
	if err != nil {
		return err
	}
	return nil
}
