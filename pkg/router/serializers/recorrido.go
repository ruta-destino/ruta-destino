package serializers

type Recorrido struct {
	Id                uint   `json:"id"`
	Dias              string `json:"dias"`
	Hora              uint   `json:"hora"`
	Minuto            uint   `json:"minuto"`
	IdEmpresa         uint   `json:"id_empresa"`
	IdTerminalOrigen  uint   `json:"id_terminal_origen"`
	IdTerminalDestino uint   `json:"id_terminal_destino"`
}
