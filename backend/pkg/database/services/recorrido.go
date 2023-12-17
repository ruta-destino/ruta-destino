package services

import (
	"ruta-destino/pkg/database/models"

	"github.com/jmoiron/sqlx"
)

type Recorrido struct {
	Db *sqlx.DB
}

func NewRecorridoService(db *sqlx.DB) *Recorrido {
	return &Recorrido{Db: db}
}

func (s *Recorrido) List(idEmpresa uint) ([]models.Recorrido, error) {
	recorridos := []models.Recorrido{}
	err := s.Db.Select(&recorridos, `
		SELECT
			recorrido.*,
			origen.nombre AS nombre_terminal_origen,
			destino.nombre AS nombre_terminal_destino
		FROM recorrido
		INNER JOIN terminal origen
		ON recorrido.id_terminal_origen = origen.id
		INNER JOIN terminal destino
		ON recorrido.id_terminal_destino = destino.id
		WHERE recorrido.id_empresa = $1
	`, idEmpresa)
	if err != nil {
		return nil, err
	}
	return recorridos, nil
}

func (s *Recorrido) Get(idEmpresa, idRecorrido uint) (*models.Recorrido, error) {
	recorrido := models.Recorrido{}
	err := s.Db.Get(&recorrido, `
		SELECT
			recorrido.*,
			origen.nombre AS nombre_terminal_origen,
			destino.nombre AS nombre_terminal_destino
		FROM recorrido
		INNER JOIN terminal origen
		ON recorrido.id_terminal_origen = origen.id
		INNER JOIN terminal destino
		ON recorrido.id_terminal_destino = destino.id
		WHERE recorrido.id_empresa = $1 AND recorrido.id = $2
	`, idEmpresa, idRecorrido)
	if err != nil {
		return nil, err
	}
	return &recorrido, nil
}

func (s *Recorrido) Insert(idEmpresa uint, recorrido *models.Recorrido) error {
	result := s.Db.QueryRow(`
		INSERT INTO recorrido (dias, hora, minuto, id_empresa, id_terminal_origen, id_terminal_destino)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`, recorrido.Dias, recorrido.Hora, recorrido.Minuto, idEmpresa, recorrido.IdTerminalOrigen, recorrido.IdTerminalDestino)
	err := result.Scan(&recorrido.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Recorrido) Update(idEmpresa, idRecorrido uint, recorrido *models.Recorrido) error {
	_, err := s.Db.Exec(`
		UPDATE recorrido
		SET dias = $1, hora = $2, minuto = $3, id_terminal_origen = $4, id_terminal_destino = $5
		WHERE id_empresa = $6 AND id = $7
	`, recorrido.Dias, recorrido.Hora, recorrido.Minuto, recorrido.IdTerminalOrigen, recorrido.IdTerminalDestino, idEmpresa, idRecorrido)
	if err != nil {
		return err
	}
	return nil
}

func (s *Recorrido) Delete(idEmpresa, idRecorrido uint) error {
	_, err := s.Db.Exec(`
		DELETE FROM recorrido
		WHERE id_empresa = $1 AND id = $2
	`, idEmpresa, idRecorrido)
	if err != nil {
		return err
	}
	return nil
}
