package serializers

type Provincia struct {
	Id           uint   `json:"id"`
	Nombre       string `json:"nombre"`
	IdRegion     uint   `json:"id_region"`
	NombreRegion string `json:"nombre_region"`
}
