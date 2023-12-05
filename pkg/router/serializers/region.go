package serializers

type Region struct {
	Id     uint   `json:"-"`
	Nombre string `json:"nombre"`
	Numero uint   `json:"numero"`
}
