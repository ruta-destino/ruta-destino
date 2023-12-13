package serializers

type Ciudad struct {
	Id          uint   `json:"id"`
	Nombre      string `json:"nombre"`
	IdProvincia uint   `json:"id_provincia"`
}
