package serializers

type Empresa struct {
	Id     uint   `json:"id"`
	Nombre string `json:"nombre"`
}

type EmpresaLinkTerminal struct {
	Id uint `json:"id_terminal"`
}
