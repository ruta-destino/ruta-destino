package models

type Recorrido struct {
	Id                uint   `db:"id"`
	Dias              string `db:"dias"`
	Hora              uint   `db:"hora"`
	Minuto            uint   `db:"minuto"`
	IdEmpresa         uint   `db:"id_empresa"`
	IdTerminalOrigen  uint   `db:"id_terminal_origen"`
	IdTerminalDestino uint   `db:"id_terminal_destino"`
}
