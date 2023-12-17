package services

import (
	"errors"
	"ruta-destino/pkg/database"
	"ruta-destino/pkg/database/models"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Empresa struct {
	Db *sqlx.DB
}

func NewEmpresaService(db *sqlx.DB) *Empresa {
	return &Empresa{Db: db}
}

func (s *Empresa) List() ([]models.Empresa, error) {
	empresas := []models.Empresa{}
	err := s.Db.Select(&empresas, `
		SELECT *
		FROM empresa
	`)
	if err != nil {
		return nil, err
	}
	return empresas, nil
}

func (s *Empresa) Get(idEmpresa uint) (*models.Empresa, error) {
	empresa := models.Empresa{}
	err := s.Db.Get(&empresa, `
		SELECT *
		FROM empresa
		WHERE id = $1
	`, idEmpresa)
	if err != nil {
		return nil, err
	}
	return &empresa, nil
}

func (s *Empresa) Insert(empresa *models.Empresa) error {
	result := s.Db.QueryRow(`
		INSERT INTO empresa (nombre)
		VALUES ($1)
		RETURNING id
	`, empresa.Nombre)
	err := result.Scan(&empresa.Id)
	if err != nil {
		if postgresError, ok := err.(*pq.Error); ok {
			return database.ProcessPostgresError(postgresError)
		}
		return errors.New("error desconocido de base de datos")
	}
	return nil
}

func (s *Empresa) Update(idEmpresa uint, empresa *models.Empresa) error {
	_, err := s.Db.Exec(`
		UPDATE empresa
		SET nombre = $1
		WHERE id = $2
	`, empresa.Nombre, idEmpresa)
	if err != nil {
		if postgresError, ok := err.(*pq.Error); ok {
			return database.ProcessPostgresError(postgresError)
		}
		return errors.New("error desconocido de base de datos")
	}
	return nil
}

func (s *Empresa) Delete(idEmpresa uint) error {
	_, err := s.Db.Exec(`
		DELETE FROM empresa
		WHERE id = $1
	`, idEmpresa)
	if err != nil {
		if postgresError, ok := err.(*pq.Error); ok {
			return database.ProcessPostgresError(postgresError)
		}
		return errors.New("error desconocido de base de datos")
	}
	return nil
}

func (s *Empresa) ListTerminales(idEmpresa uint) ([]models.Terminal, error) {
	terminales := []models.Terminal{}
	err := s.Db.Select(&terminales, `
		SELECT terminal.*
		FROM terminal
		INNER JOIN empresa_terminal
		ON terminal.id = empresa_terminal.id_terminal
		WHERE empresa_terminal.id_empresa = $1
	`, idEmpresa)
	if err != nil {
		return nil, err
	}
	return terminales, nil
}

func (s *Empresa) LinkTerminal(idEmpresa, idTerminal uint) error {
	_, err := s.Db.Exec(`
		INSERT INTO empresa_terminal (id_empresa, id_terminal)
		VALUES ($1, $2)
	`, idEmpresa, idTerminal)
	if err != nil {
		if postgresError, ok := err.(*pq.Error); ok {
			return database.ProcessPostgresError(postgresError)
		}
		return errors.New("error desconocido de base de datos")
	}
	return nil
}

func (s *Empresa) UnlinkTerminal(idEmpresa, idTerminal uint) error {
	_, err := s.Db.Exec(`
		DELETE FROM empresa_terminal
		WHERE id_empresa = $1 AND id_terminal = $2
	`, idEmpresa, idTerminal)
	if err != nil {
		if postgresError, ok := err.(*pq.Error); ok {
			return database.ProcessPostgresError(postgresError)
		}
		return errors.New("error desconocido de base de datos")
	}
	return nil
}
