package services

import (
	"errors"
	"ruta-destino/pkg/database"
	"ruta-destino/pkg/database/models"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Terminal struct {
	Db *sqlx.DB
}

func NewTerminalService(db *sqlx.DB) *Terminal {
	return &Terminal{Db: db}
}

func (s *Terminal) List() ([]models.Terminal, error) {
	terminales := []models.Terminal{}
	err := s.Db.Select(&terminales, `
		SELECT terminal.*, ciudad.nombre AS "nombre_ciudad"
		FROM terminal
		INNER JOIN ciudad
		ON terminal.id_ciudad = ciudad.id
	`)
	if err != nil {
		return nil, err
	}
	return terminales, nil
}

func (s *Terminal) Get(idTerminal uint) (*models.Terminal, error) {
	terminal := models.Terminal{}
	err := s.Db.Get(&terminal, `
		SELECT *
		FROM terminal
		WHERE id = $1
	`, idTerminal)
	if err != nil {
		return nil, err
	}
	return &terminal, nil
}

func (s *Terminal) Insert(terminal *models.Terminal) error {
	result := s.Db.QueryRow(`
		INSERT INTO terminal (nombre, longitud, latitud, direccion, id_ciudad)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`, terminal.Nombre, terminal.Longitud, terminal.Latitud, terminal.Direccion, terminal.IdCiudad)
	err := result.Scan(&terminal.Id)
	if err != nil {
		if postgresError, ok := err.(*pq.Error); ok {
			return database.ProcessPostgresError(postgresError)
		}
		return errors.New("error desconocido de base de datos")
	}
	return nil
}

func (s *Terminal) Update(idTerminal uint, terminal *models.Terminal) error {
	_, err := s.Db.Exec(`
		UPDATE terminal
		SET nombre = $1, longitud = $2, latitud = $3, direccion = $4, id_ciudad = $5
		WHERE id = $6
	`, terminal.Nombre, terminal.Longitud, terminal.Latitud, terminal.Direccion, terminal.IdCiudad, idTerminal)
	if err != nil {
		if postgresError, ok := err.(*pq.Error); ok {
			return database.ProcessPostgresError(postgresError)
		}
		return errors.New("error desconocido de base de datos")
	}
	return nil
}

func (s *Terminal) Delete(idTerminal uint) error {
	_, err := s.Db.Exec(`
		DELETE FROM terminal
		WHERE id = $1
	`, idTerminal)
	if err != nil {
		if postgresError, ok := err.(*pq.Error); ok {
			return database.ProcessPostgresError(postgresError)
		}
		return errors.New("error desconocido de base de datos")
	}
	return nil
}

func (s *Terminal) ListEmpresa(idTerminal uint) ([]models.Empresa, error) {
	empresas := []models.Empresa{}
	err := s.Db.Select(&empresas, `
		SELECT empresa.*
		FROM empresa
		INNER JOIN empresa_terminal
		ON empresa.id = empresa_terminal.id_empresa
		WHERE empresa_terminal.id_terminal = $1
	`, idTerminal)
	if err != nil {
		return nil, err
	}
	return empresas, nil
}
